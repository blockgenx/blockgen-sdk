package main

import (
	"os"

	"github.com/blockgenx/blockgen-sdk/server"
	svrcmd "github.com/blockgenx/blockgen-sdk/server/cmd"
	"github.com/blockgenx/blockgen-sdk/simapp"
	"github.com/blockgenx/blockgen-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
