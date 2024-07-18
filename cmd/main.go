package main

import (
	"GoRestBlockchain/pkg/blockchain"
	cli2 "GoRestBlockchain/pkg/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	var cli = cli2.CLI{bc}
	cli.Run()

}
