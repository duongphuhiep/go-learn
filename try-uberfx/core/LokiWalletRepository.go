package core

import "log"

type LokiWalletRepository struct {
	name string
}

var _ IWalletRepository = (*LokiWalletRepository)(nil)

func NewLokiWalletRepository() *LokiWalletRepository {
	log.Println("NewLokiWalletRepository is called")
	return &LokiWalletRepository{name: "LokiWallet"}
}

func (this *LokiWalletRepository) GetWalletNameFromDb() string {
	return this.name
}
