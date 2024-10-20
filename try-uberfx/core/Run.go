package core

import "log"

func Run(g IWalletService) {
	log.Println("Run is called 1")
	log.Println(g.GetWalletDetails())

	log.Println("Run is called 2")
	log.Println(g.GetWalletDetails())
}
