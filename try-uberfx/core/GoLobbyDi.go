package core

import (
	"log"

	"github.com/golobby/container/v3"
)

func Setup() {
	err := container.Singleton(func() IWalletRepository {
		return NewLokiWalletRepository()
	})
	if err != nil {
		panic(err)
	}
	err = container.Singleton(func() IWalletRepository {
		return NewThorWalletRepository()
	})
	if err != nil {
		panic(err)
	}
	var repo []IWalletRepository
	err = container.Resolve(&repo)
	if err != nil {
		panic(err)
	}
	log.Println(len(repo))
}
