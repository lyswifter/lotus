package main

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/urfave/cli/v2"
)

func GetLocalWalletApi(cctx *cli.Context) (api.Wallet, error) {
	lr, ks, err := openRepo(cctx)
	if err != nil {
		return nil, err
	}
	defer lr.Close() // nolint

	lw, err := wallet.NewWallet(ks)
	if err != nil {
		return nil, err
	}

	var w api.Wallet = lw

	return w, nil
}
