package main

import (
	"log"
	"time"
	"github.com/ayusrina/todo2/storage"
	"github.com/ayusrina/todo2/model"

)

func main() {
	//init dan memilih tipe storage
	var memStore = storage.GetStorage(storage.StorageTypeDatabase)

	//create
	if err := memStore.Create(model.Todo{
		ID: 	1,
		Title: 	"first",
		Description: "First Todo",
		CreatedAt: 	time.Now(),
	
	}); err != nil {
		log.Fatal(err)
	}

	// if err := memStore.Create(model.Todo{
	// 	ID: 	2,
	// 	Title:	"second",
	// 	Description: "second todo",
	// 	CreatedAt: time.Now(),
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("todo created")

	//detail
	obj, err := memStore.Detail(1)
	if err != nil{
		log.Fatal(err)
	
	}
	log.Printf("%d: %s\n", obj.ID,obj.Title)

	//list
	list, err :=  memStore.List()
	if err != nil {
		log.Fatal (err)
	} 
	for _, res := range list{
		log.Printf("ID: %d, Title: %s, Description: %s \n", res.ID, res.Title, res.Description)
		
	}
	
}



