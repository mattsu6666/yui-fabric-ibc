package main

import (
	"fmt"
	"io"

	"github.com/datachainlab/fabric-ibc/app"
	"github.com/datachainlab/fabric-ibc/chaincode"
	"github.com/datachainlab/fabric-ibc/example"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func main() {
	cc := chaincode.NewIBCChaincode(newApp, chaincode.DefaultDBProvider)
	chaincode, err := contractapi.NewChaincode(cc)

	if err != nil {
		fmt.Printf("Error create IBC chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting IBC chaincode: %s", err.Error())
		return
	}
}

func newApp(logger tmlog.Logger, db tmdb.DB, traceStore io.Writer, blockProvider app.BlockProvider) (app.Application, error) {
	return example.NewIBCApp(
		logger,
		db,
		traceStore,
		example.MakeEncodingConfig(),
		blockProvider,
	)
}
