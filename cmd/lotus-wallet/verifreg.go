package main

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	lcli "github.com/filecoin-project/lotus/cli"
	builtin8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"
	"github.com/urfave/cli/v2"

	vrctor1 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
)

var verifyClientCmd = &cli.Command{
	Name:  "verify",
	Usage: "Interact with a verify client",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "confidence",
			Usage: "number of block confirmations to wait for",
			Value: int(build.MessageConfidence),
		},
	},
	Subcommands: []*cli.Command{
		addClientCmd,
	},
}

var addClientCmd = &cli.Command{
	Name:  "add",
	Usage: "Add verify client",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "the address used to sign the message",
		},
		&cli.StringFlag{
			Name:  "client",
			Usage: "client address to specify",
		},
		&cli.StringFlag{
			Name:  "allowance",
			Usage: "client allowance to specify",
		},
	},
	Action: func(cctx *cli.Context) error {
		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		lw, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}
		var lwapi api.Wallet = lw

		srv, err := lcli.GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		fapi := srv.FullNodeAPI()

		ctx := lcli.ReqContext(cctx)

		from, err := address.NewFromString(cctx.String("from"))
		if err != nil {
			return err
		}

		fromact, err := fapi.StateGetActor(ctx, from, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to look up address %s: %w", from, err)
		}

		client, err := address.NewFromString(cctx.String("client"))
		if err != nil {
			return err
		}

		allowance, err := big.FromString(cctx.String("allowance"))
		if err != nil {
			return err
		}

		var params = &vrctor1.AddVerifiedClientParams{
			Address:   client,
			Allowance: allowance,
		}

		buf, err := actors.SerializeParams(params)
		if err != nil {
			return err
		}

		dmsg := &types.Message{
			To:     builtin8.VerifiedRegistryActorAddr,
			From:   from,
			Value:  big.Zero(),
			Method: builtin8.MethodsVerifiedRegistry.AddVerifiedClient,
			Params: buf,
		}

		gasedMsg, err := fapi.GasEstimateMessageGas(ctx, dmsg, nil, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("estimating gas: %w", err)
		}
		dmsg = gasedMsg
		dmsg.Nonce = fromact.Nonce

		keyAddr, err := fapi.StateAccountKey(ctx, from, types.EmptyTSK)
		if err != nil {
			return fmt.Errorf("failed to resolve ID address: %s", keyAddr.String())
		}

		mb, err := dmsg.ToStorageBlock()
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
			Message:   *dmsg,
			Signature: *sig,
		})
		if err != nil {
			return err
		}

		msgCid := cid

		fmt.Println("sent add verify client in message: ", msgCid)

		wait, err := fapi.StateWaitMsg(ctx, msgCid, uint64(cctx.Int("confidence")), build.Finality, true)
		if err != nil {
			return err
		}

		if wait.Receipt.ExitCode != 0 {
			return fmt.Errorf("proposal returned exit %d", wait.Receipt.ExitCode)
		}

		fmt.Println("Add verify client finished")

		return nil
	},
}
