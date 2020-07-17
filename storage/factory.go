package storage

import "log"

type  StorageType int

const (
	StorageTypeMemory StorageType = iota
	StorageTypeMock
	StorageTypeDatabase
)

func GetStorage(t StorageType) Storage {
	// memilih tipe data storage berdasarkan struct 
	var s Storage
	switch t {
	case StorageTypeMemory: // kalau tipe memory, jalanin struct nya memory
		s = newMemory()
	case StorageTypeMock: // kalau tipe mock, jalanin struct nya mock
		s = Mock{}
	case StorageTypeDatabase: // kalau tipe database, jalanin struct nya database
	 	s = Database{}	
	default:
		log.Fatal("Storage type tidak diketahui")
	}
	return s 
}