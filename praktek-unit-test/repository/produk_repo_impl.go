package repository

import (
	"praktek-unit-test/entity"
	"github.com/stretchr/testify/mock"
)
type RepositoryProdukService struct{
	Mock mock.Mock
}

func (repository RepositoryProdukService)GetByNama(nama string) *entity.Produk{
	argument := repository.Mock.Called(nama)
	if argument.Get(0) == nil{
		return nil
	}else{
		produk := argument.Get(0).(*entity.Produk)
		return produk
	}
}