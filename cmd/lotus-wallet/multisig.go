package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	lcli "github.com/filecoin-project/lotus/cli"
	msig2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var multisigCmd = &cli.Command{
	Name:  "msig",
	Usage: "Interact with a multisig wallet",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "confidence",
			Usage: "number of block confirmations to wait for",
			Value: int(build.MessageConfidence),
		},
	},
	Subcommands: []*cli.Command{
		msigProposeCmd,
		msigApproveCmd,
	},
}

var msigProposeCmd = &cli.Command{
	Name:      "propose",
	Usage:     "Propose a multisig transaction",
	ArgsUsage: "[multisigAddress destinationAddress value <methodId methodParams> (optional)]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "account to send the propose message from",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 3 {
			return lcli.ShowHelp(cctx, fmt.Errorf("must pass at least multisig address, destination, and value"))
		}

		if cctx.Args().Len() > 3 && cctx.Args().Len() != 5 {
			return lcli.ShowHelp(cctx, fmt.Errorf("must either pass three or five arguments"))
		}

		lwapi, err := GetLocalWalletApi(cctx)
		if err != nil {
			return err
		}

		srv, err := lcli.GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		fapi := srv.FullNodeAPI()

		ctx := lcli.ReqContext(cctx)

		msig, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		dest, err := address.NewFromString(cctx.Args().Get(1))
		if err != nil {
			return err
		}

		value, err := types.ParseFIL(cctx.Args().Get(2))
		if err != nil {
			return err
		}

		var method uint64
		var params []byte
		if cctx.Args().Len() == 5 {
			m, err := strconv.ParseUint(cctx.Args().Get(3), 10, 64)
			if err != nil {
				return err
			}
			method = m

			p, err := hex.DecodeString(cctx.Args().Get(4))
			if err != nil {
				return err
			}
			params = p
		}

		var from address.Address
		if cctx.IsSet("from") {
			f, err := address.NewFromString(cctx.String("from"))
			if err != nil {
				return err
			}
			from = f
		} else {
			lr, ks, err := openRepo(cctx)
			if err != nil {
				return err
			}
			defer lr.Close() // nolint

			ki, err := ks.Get(wallet.KDefault)
			if err != nil {
				return fmt.Errorf("failed to get default key: %w", err)
			}

			k, err := key.NewKey(ki)
			if err != nil {
				return fmt.Errorf("failed to read default key from keystore: %w", err)
			}

			from = k.Address
		}

		act, err := fapi.StateGetActor(ctx, msig, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to look up multisig %s: %w", msig, err)
		}

		if !builtin.IsMultisigActor(act.Code) {
			return fmt.Errorf("actor %s is not a multisig actor", msig)
		}

		curNonce := act.Nonce

		proto, err := fapi.MsigPropose(ctx, msig, dest, types.BigInt(value), from, method, params)
		if err != nil {
			return err
		}

		gasedMsg, err := fapi.GasEstimateMessageGas(ctx, &proto.Message, nil, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("estimating gas: %w", err)
		}
		proto.Message = *gasedMsg
		proto.Message.Nonce = curNonce + 1

		keyAddr, err := fapi.StateAccountKey(ctx, proto.Message.From, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to resolve ID address: %s", keyAddr.String())
		}

		mb, err := proto.Message.ToStorageBlock()
		if err != nil {
			return fmt.Errorf("serializing message: %w", err)
		}

		sig, err := lwapi.WalletSign(ctx, keyAddr, mb.Cid().Bytes(), api.MsgMeta{
			Type:  api.MTChainMsg,
			Extra: mb.RawData(),
		})
		if err != nil {
			return fmt.Errorf("failed to sign message: %w", err)
		}

		cid, err := fapi.MpoolPush(ctx, &types.SignedMessage{
			Message:   proto.Message,
			Signature: *sig,
		})
		if err != nil {
			return err
		}

		// sm, err := lcli.InteractiveSend(ctx, cctx, srv, proto)
		// if err != nil {
		// 	return err
		// }

		msgCid := cid

		fmt.Println("sent proposal in message: ", msgCid)

		wait, err := fapi.StateWaitMsg(ctx, msgCid, uint64(cctx.Int("confidence")), build.Finality, true)
		if err != nil {
			return err
		}

		if wait.Receipt.ExitCode != 0 {
			return fmt.Errorf("proposal returned exit %d", wait.Receipt.ExitCode)
		}

		var retval msig2.ProposeReturn
		if err := retval.UnmarshalCBOR(bytes.NewReader(wait.Receipt.Return)); err != nil {
			return fmt.Errorf("failed to unmarshal propose return value: %w", err)
		}

		fmt.Printf("Transaction ID: %d\n", retval.TxnID)
		if retval.Applied {
			fmt.Printf("Transaction was executed during propose\n")
			fmt.Printf("Exit Code: %d\n", retval.Code)
			fmt.Printf("Return Value: %x\n", retval.Ret)
		}

		return nil
	},
}

var msigApproveCmd = &cli.Command{
	Name:      "approve",
	Usage:     "Approve a multisig message",
	ArgsUsage: "<multisigAddress messageId> [proposerAddress destination value [methodId methodParams]]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "account to send the approve message from",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 2 {
			return lcli.ShowHelp(cctx, fmt.Errorf("must pass at least multisig address and message ID"))
		}

		if cctx.Args().Len() > 2 && cctx.Args().Len() < 5 {
			return lcli.ShowHelp(cctx, fmt.Errorf("usage: msig approve <msig addr> <message ID> <proposer address> <desination> <value>"))
		}

		if cctx.Args().Len() > 5 && cctx.Args().Len() != 7 {
			return lcli.ShowHelp(cctx, fmt.Errorf("usage: msig approve <msig addr> <message ID> <proposer address> <desination> <value> [ <method> <params> ]"))
		}

		lwapi, err := GetLocalWalletApi(cctx)
		if err != nil {
			return err
		}

		srv, err := lcli.GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		fapi := srv.FullNodeAPI()
		ctx := lcli.ReqContext(cctx)

		msig, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		txid, err := strconv.ParseUint(cctx.Args().Get(1), 10, 64)
		if err != nil {
			return err
		}

		var from address.Address
		if cctx.IsSet("from") {
			f, err := address.NewFromString(cctx.String("from"))
			if err != nil {
				return err
			}
			from = f
		} else {
			lr, ks, err := openRepo(cctx)
			if err != nil {
				return err
			}
			defer lr.Close() // nolint

			ki, err := ks.Get(wallet.KDefault)
			if err != nil {
				return fmt.Errorf("failed to get default key: %w", err)
			}

			k, err := key.NewKey(ki)
			if err != nil {
				return fmt.Errorf("failed to read default key from keystore: %w", err)
			}

			from = k.Address
		}

		act, err := fapi.StateGetActor(ctx, from, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to look up multisig %s: %w", msig, err)
		}

		curNonce := act.Nonce

		var msgCid cid.Cid
		if cctx.Args().Len() == 2 {
			proto, err := fapi.MsigApprove(ctx, msig, txid, from)
			if err != nil {
				return err
			}

			gasedMsg, err := fapi.GasEstimateMessageGas(ctx, &proto.Message, nil, types.EmptyTSK)
			if err != nil {
				return fmt.Errorf("estimating gas: %w", err)
			}
			proto.Message = *gasedMsg
			proto.Message.Nonce = curNonce + 1

			keyAddr, err := fapi.StateAccountKey(ctx, proto.Message.From, types.EmptyTSK)
			if err != nil {
				return fmt.Errorf("failed to resolve ID address: %s", keyAddr.String())
			}

			mb, err := proto.Message.ToStorageBlock()
			if err != nil {
				return fmt.Errorf("serializing message: %w", err)
			}

			sig, err := lwapi.WalletSign(ctx, keyAddr, mb.Cid().Bytes(), api.MsgMeta{
				Type:  api.MTChainMsg,
				Extra: mb.RawData(),
			})
			if err != nil {
				return fmt.Errorf("failed to sign message: %w", err)
			}

			cid, err := fapi.MpoolPush(ctx, &types.SignedMessage{
				Message:   proto.Message,
				Signature: *sig,
			})
			if err != nil {
				return err
			}

			// sm, err := lcli.InteractiveSend(ctx, cctx, srv, proto)
			// if err != nil {
			// 	return err
			// }

			msgCid = cid
		} else {
			proposer, err := address.NewFromString(cctx.Args().Get(2))
			if err != nil {
				return err
			}

			if proposer.Protocol() != address.ID {
				proposer, err = fapi.StateLookupID(ctx, proposer, types.EmptyTSK)
				if err != nil {
					return err
				}
			}

			dest, err := address.NewFromString(cctx.Args().Get(3))
			if err != nil {
				return err
			}

			value, err := types.ParseFIL(cctx.Args().Get(4))
			if err != nil {
				return err
			}

			var method uint64
			var params []byte
			if cctx.Args().Len() == 7 {
				m, err := strconv.ParseUint(cctx.Args().Get(5), 10, 64)
				if err != nil {
					return err
				}
				method = m

				p, err := hex.DecodeString(cctx.Args().Get(6))
				if err != nil {
					return err
				}
				params = p
			}

			proto, err := fapi.MsigApproveTxnHash(ctx, msig, txid, proposer, dest, types.BigInt(value), from, method, params)
			if err != nil {
				return err
			}

			gasedMsg, err := fapi.GasEstimateMessageGas(ctx, &proto.Message, nil, types.EmptyTSK)
			if err != nil {
				return fmt.Errorf("estimating gas: %w", err)
			}
			proto.Message = *gasedMsg
			proto.Message.Nonce = curNonce + 1

			keyAddr, err := fapi.StateAccountKey(ctx, proto.Message.From, types.EmptyTSK)
			if err != nil {
				return fmt.Errorf("failed to resolve ID address: %s", keyAddr.String())
			}

			mb, err := proto.Message.ToStorageBlock()
			if err != nil {
				return fmt.Errorf("serializing message: %w", err)
			}

			sig, err := lwapi.WalletSign(ctx, keyAddr, mb.Cid().Bytes(), api.MsgMeta{
				Type:  api.MTChainMsg,
				Extra: mb.RawData(),
			})
			if err != nil {
				return fmt.Errorf("failed to sign message: %w", err)
			}

			cid, err := fapi.MpoolPush(ctx, &types.SignedMessage{
				Message:   proto.Message,
				Signature: *sig,
			})
			if err != nil {
				return err
			}

			// sm, err := lcli.InteractiveSend(ctx, cctx, srv, proto)
			// if err != nil {
			// 	return err
			// }

			msgCid = cid
		}

		fmt.Println("sent approval in message: ", msgCid)

		wait, err := fapi.StateWaitMsg(ctx, msgCid, uint64(cctx.Int("confidence")), build.Finality, true)
		if err != nil {
			return err
		}

		if wait.Receipt.ExitCode != 0 {
			return fmt.Errorf("approve returned exit %d", wait.Receipt.ExitCode)
		}

		return nil
	},
}
