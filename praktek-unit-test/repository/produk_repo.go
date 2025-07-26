package repository
import "praktek-unit-test/entity"
type RepositoryProduk interface {
	GetByNama(nama string) *entity.Produk
}