package core

import (
	"fmt"
	"log"
)

var _ IWalletService = (*WalletServiceImpl)(nil)

type WalletServiceImpl struct {
	walletRepostiory IWalletRepository
}

func NewWalletServiceImpl(walletRepostiory IWalletRepository) *WalletServiceImpl {
	log.Println("NewWalletServiceImpl is called")
	return &WalletServiceImpl{walletRepostiory}
}

func (this *WalletServiceImpl) GetWalletDetails() string {
	return fmt.Sprintf("Wallet details is %s!", this.walletRepostiory.GetWalletNameFromDb())
}
