package service

import (
	"praktek-unit-test/entity"
	"praktek-unit-test/repository"
	"errors"
)

type ServiceRepository struct{
	Repository repository.RepositoryProduk
}

func (service ServiceRepository)GetProduk(nama string) (*entity.Produk,error){
	produk := service.Repository.GetByNama(nama)
	if produk == nil{
		return nil,errors.New("tidak ditemukan produk")
	}else{
		return produk,nil
	}
}