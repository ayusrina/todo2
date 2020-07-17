package storage

import (
	"errors"
	"time"
	"github.com/ayusrina/todo2/model"

)

// membuat objek dengan nama Mock
type Mock struct{}

func (o Mock) Create(obj model.Todo) error {
	//langsung return kosong karena mock
	return nil

}

func (o Mock) Detail(id int32) (model.Todo, error) {
	// kalau id nya 1 langsung return hasil yang sudah di hardcode
	// selain 1 langsung return tidak ditemukan
	if id == 1 {
		return model.Todo{ID : 1, Title: "Mock Title", Description: "Mock Description", CreatedAt: time.Now()}, nil
	}
	return model.Todo{}, errors.New("todo tidak ditemukan")
}

func (o Mock) Delete(id int32) error {
	//langsung return kosong karena mock
	return nil
}

func (o Mock) List() ([]model.Todo, error) {
	list := []model.Todo{} // inisialisasi variabel list dengan array model
	// isi list dengan data yang di sudah di set
	list = append(list, model.Todo {
		ID: 	1,
		Title: 	"firstMock",
		Description: "First Todo mock",
		CreatedAt: 	time.Now(),
	})
	// , model.Todo{

	// 	ID: 	2,
	// 	Title:	"secondMock",
	// 	Description: "second todo mock",
	// 	CreatedAt: time.Now(),			
	// }
	
	// return list nya dengan error null
	return list, nil
}