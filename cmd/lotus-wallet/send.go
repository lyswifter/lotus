package main

import (
	"encoding/hex"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	lcli "github.com/filecoin-project/lotus/cli"

	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
)

var sendCmd = &cli.Command{
	Name:      "send",
	Usage:     "Send funds between accounts",
	ArgsUsage: "[targetAddress] [amount]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "optionally specify the account to send funds from",
		},
		&cli.StringFlag{
			Name:  "gas-premium",
			Usage: "specify gas price to use in AttoFIL",
			Value: "0",
		},
		&cli.StringFlag{
			Name:  "gas-feecap",
			Usage: "specify gas fee cap to use in AttoFIL",
			Value: "0",
		},
		&cli.Int64Flag{
			Name:  "gas-limit",
			Usage: "specify gas limit",
			Value: 0,
		},
		&cli.Uint64Flag{
			Name:  "nonce",
			Usage: "specify the nonce to use",
			Value: 0,
		},
		&cli.Uint64Flag{
			Name:  "method",
			Usage: "specify method to invoke",
			Value: uint64(builtin.MethodSend),
		},
		&cli.StringFlag{
			Name:  "params-json",
			Usage: "specify invocation parameters in json",
		},
		&cli.StringFlag{
			Name:  "params-hex",
			Usage: "specify invocation parameters in hex",
		},
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Deprecated: use global 'force-send'",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet("force") {
			fmt.Println("'force' flag is deprecated, use global flag 'force-send'")
		}

		if cctx.Args().Len() != 2 {
			return lcli.ShowHelp(cctx, fmt.Errorf("'send' expects two arguments, target and amount"))
		}

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		lw, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		var wapi api.Wallet = lw

		srv, err := lcli.GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		fapi := srv.FullNodeAPI()

		ctx := lcli.ReqContext(cctx)
		var params lcli.SendParams

		params.To, err = address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return lcli.ShowHelp(cctx, fmt.Errorf("failed to parse target address: %w", err))
		}

		val, err := types.ParseFIL(cctx.Args().Get(1))
		if err != nil {
			return lcli.ShowHelp(cctx, fmt.Errorf("failed to parse amount: %w", err))
		}
		params.Val = abi.TokenAmount(val)

		if from := cctx.String("from"); from != "" {
			addr, err := address.NewFromString(from)
			if err != nil {
				return err
			}

			params.From = addr
		}

		if cctx.IsSet("gas-premium") {
			gp, err := types.BigFromString(cctx.String("gas-premium"))
			if err != nil {
				return err
			}
			params.GasPremium = &gp
		}

		if cctx.IsSet("gas-feecap") {
			gfc, err := types.BigFromString(cctx.String("gas-feecap"))
			if err != nil {
				return err
			}
			params.GasFeeCap = &gfc
		}

		if cctx.IsSet("gas-limit") {
			limit := cctx.Int64("gas-limit")
			params.GasLimit = &limit
		}

		params.Method = abi.MethodNum(cctx.Uint64("method"))

		if cctx.IsSet("params-json") {
			decparams, err := srv.DecodeTypedParamsFromJSON(ctx, params.To, params.Method, cctx.String("params-json"))
			if err != nil {
				return fmt.Errorf("failed to decode json params: %w", err)
			}
			params.Params = decparams
		}
		if cctx.IsSet("params-hex") {
			if params.Params != nil {
				return fmt.Errorf("can only specify one of 'params-json' and 'params-hex'")
			}
			decparams, err := hex.DecodeString(cctx.String("params-hex"))
			if err != nil {
				return fmt.Errorf("failed to decode hex params: %w", err)
			}
			params.Params = decparams
		}

		if cctx.IsSet("nonce") {
			n := cctx.Uint64("nonce")
			params.Nonce = &n
		}

		act, err := fapi.StateGetActor(ctx, params.From, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to look up multisig %s: %w", params.From, err)
		}

		proto, err := srv.MessageForSend(ctx, params)
		if err != nil {
			return fmt.Errorf("creating message prototype: %w", err)
		}

		gasedMsg, err := fapi.GasEstimateMessageGas(ctx, &proto.Message, nil, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("estimating gas: %w", err)
		}

		fmt.Printf("act.Nonce %d", act.Nonce)

		proto.Message = *gasedMsg
		proto.Message.Nonce = act.Nonce

		keyAddr, err := fapi.StateAccountKey(ctx, proto.Message.From, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to resolve ID address: %s", keyAddr.String())
		}

		mb, err := proto.Message.ToStorageBlock()
		if err != nil {
			return fmt.Errorf("serializing message: %w", err)
		}

		sig, err := wapi.WalletSign(ctx, keyAddr, mb.Cid().Bytes(), api.MsgMeta{
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

		msgCid := cid

		// sm, err := lcli.InteractiveSend(ctx, cctx, srv, proto)
		// if err != nil {
		// 	return err
		// }

		fmt.Fprintf(cctx.App.Writer, "%s\n", msgCid)
		fmt.Println("waiting for confirmation..")

		// wait for it to get mined into a block
		wait, err := fapi.StateWaitMsg(ctx, msgCid, uint64(cctx.Int("confidence")), build.Finality, true)
		if err != nil {
			return err
		}

		// check it executed successfully
		if wait.Receipt.ExitCode != 0 {
			fmt.Fprintln(cctx.App.Writer, "actor creation failed!")
			return err
		}
		fmt.Fprintln(cctx.App.Writer, "Send fil finished")

		return nil
	},
}
