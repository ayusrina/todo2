package storage

import (
	"github.com/ayusrina/todo2/model"
	"database/sql"
	//library postgres
	_"github.com/lib/pq"
	"log"
	"time"
)

type Database struct{} 

func Connection() (db *sql.DB, err error) {
	// open koneksi postgres nya
	db, err = sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres " + "password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	return db, err 
}

func (o Database) Create(obj model.Todo) error{
	// panggil function Connection supaya terkoneksi sama postgres
	db, _ := Connection()
	// execute query insert
	// $ --> cara define variable di postgres
	_, err := db.Exec("INSERT INTO tododb (id, title, description, created) VALUES " +
	"($1,$2,$3,$4);", obj.ID, obj.Title, obj.Description, time.Now())

	// kalau ada error, di return errornya, kalau sukses ga return apa2
	return err

}

func (o Database) Detail(id int32) (model.Todo, error){
	// panggil function Connection supaya terkoneksi sama postgres
	db, _ := Connection()
	//execute query select
	rows, err := db.Query("SELECT id, title, description, created FROM tododb WHERE id = $1;",id)
	// kalau ada errornya langsung return hasilnya null, error nya ada (skip if nya kalau ga error)
	if err != nil {
		return model.Todo{}, err
	}

	defer rows.Close()
	// looping selama masih ada row berikutnya
	for rows.Next() {
		var res model.Todo
		err := rows.Scan(&res.ID, &res.Title, &res.Description, &res.CreatedAt)

		if err !=nil {
			return model.Todo{}, err
		}
		// return hasil loop nya dengan format ID:, Title: , dll
		return model.Todo{ID: res.ID, Title: res.Title, Description: res.Description, CreatedAt: time.Now()}, nil
	}
	// return hasil akhir tanpa error 
	return model.Todo{}, nil	
}

func (o Database) Delete(id int32) error{
	db, _ := Connection()
	//execute query delete
	_, err := db.Exec("DELETE from tododb WHERE id = $1",id)
	return err
}


func (o Database) List() ([]model.Todo, error) {
	db, _ := Connection()

	rows, err := db.Query("SELECT id, title, description, created FROM tododb LIMIT 10;")
	// kalau ada errornya langsung return hasilnya null, error nya ada (skip if nya kalau ga error)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make([]model.Todo, 0) //inisialisasi variabel list dengan array model dan size 0

	for rows.Next() {
		var res model.Todo
		err := rows.Scan(&res.ID, &res.Title, &res.Description, &res.CreatedAt)

		// kalau ada errornya langsung return hasilnya null, error nya ada (skip if nya kalau ga error)
		if err != nil {
			return nil, err
		}
		// tambahkan hasil looping data nya ke list
		list = append (list,res)
	}
	// kalau ga error return data nya dengan error nil
	return list, nil 
}
