package storage

import "github.com/ayusrina/todo2/model"

type Storage interface {
	Create(obj model.Todo) error // menyimpan objek model
	Detail(id int32) (model.Todo, error) // mengambil objek model sesuai id
	Delete(id int32) error // menghapus objek model sesuai id
	List()([]model.Todo, error) // mengambil seluruh data objek model
}