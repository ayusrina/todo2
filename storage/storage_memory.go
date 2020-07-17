package storage

import (
	"github.com/ayusrina/todo2/model"
	"errors"
)

// membuat objek dengan nama Memory
type Memory struct {
	store map[int32]model.Todo
}

func newMemory() Memory {
	return Memory{
		store:make(map[int32]model.Todo),
	}
}

func (o Memory) Create(obj model.Todo) error{
	// store data di memory
	o.store[obj.ID] = obj
	return nil
}

func (o Memory) Detail(id int32) (model.Todo, error){
	// cek apakah ada data sama dengan id yang dicari
	if obj, ok := o.store[id]; ok {
		return obj, nil // kalau ada return datanya
	}
	return model.Todo{}, errors.New("todo tidak ditemukan") // kalau ga ada return error data tidak ditemukan
}

func (o Memory) Delete(id int32) error{
	// delete data di memory dengan id inputan
	delete(o.store, id)
	return nil
}

func (o Memory) List() ([]model.Todo, error) {
	list := []model.Todo{} // inisialisasi variabel list dengan array model
	// isi list dengan looping data yang ada di memory
	for _, res := range o.store {
		list = append (list, res) // di append untuk mengisi list
	}
	return list, nil
}