package main

import (
	"cddstore/app"
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	tmos "github.com/tendermint/tendermint/libs/os"
	"log"
	"os"
)

func main() {
	cfg := app.InitConfig()
	db := app.CreateDB(cfg)
	app, err := app.NewStoreApp(db)
	if err != nil {
		fmt.Println("new app", err)
		os.Exit(1)
	}
	svr, err := server.NewServer("127.0.0.1:26658", "socket", app)
	if err != nil {
		fmt.Println("new server", err)
		os.Exit(2)
	}

	err = svr.Start()
	if err != nil {
		log.Println("start server", err)
		os.Exit(3)
	}

	tmos.TrapSignal(nil, func() {
		svr.Stop()
	})

	select {

	}
}
