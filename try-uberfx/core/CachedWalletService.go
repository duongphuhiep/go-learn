package core

import (
	"fmt"
	"log"
)

var _ IWalletService = (*CachedWalletService)(nil)

type CachedWalletService struct {
	_core IWalletService
}

func NewCachedWalletService(walletRepostiory IWalletService) *CachedWalletService {
	log.Println("NewCachedWalletService is called")
	return &CachedWalletService{walletRepostiory}
}

func (this *CachedWalletService) GetWalletDetails() string {
	log.Println("Get wallet details from cache")
	return fmt.Sprintf("Wallet details is %s!", this._core.GetWalletDetails())
}
