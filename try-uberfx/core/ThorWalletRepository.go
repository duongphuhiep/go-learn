package core

import "log"

type ThorWalletRepository struct {
	name string
}

var _ IWalletRepository = (*ThorWalletRepository)(nil)

func NewThorWalletRepository() *ThorWalletRepository {
	log.Println("NewThorWalletRepository is called")
	return &ThorWalletRepository{name: "ThorWallet"}
}

func (this *ThorWalletRepository) GetWalletNameFromDb() string {
	return this.name
}
