package service

import (
	"praktek-unit-test/entity"
	"praktek-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
var RepositoryMockService = &repository.RepositoryProdukService{Mock: mock.Mock{},}
var RepositoryService = ServiceRepository{Repository: RepositoryMockService,}

func TestProduk(t *testing.T){
	data := []struct{
		nama string
		request string
		expected *entity.Produk
	}{
		{"produk 1","roti",&entity.Produk{Nama : "roti", ID : 1, Harga : 5000}},
		{"produk 2","keju",&entity.Produk{Nama : "keju", ID : 2, Harga : 5000}},
		{"produk 3","berondong",nil},
	}
	for _,field := range data {
		t.Run(field.nama,func(t *testing.T){
			RepositoryMockService.Mock.On("GetByNama",field.request).Return(field.expected)
			hasil,err := RepositoryService.GetProduk(field.request)

			if field.expected == nil {
				assert.Nil(t, hasil)
				assert.NotNil(t, err)
			} else {
				assert.NotNil(t, hasil)
				assert.Nil(t, err)
				assert.Equal(t, field.expected, hasil)
			}

		})
	}
}