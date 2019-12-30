package core

import (
	"fmt"
	"github.com/boltdb/bolt"
)

var (
	// local persistent storage bucket name
	gokv_store = []byte("gokv_sotre")
)

type PersistentStorage struct {
	// boltdb as persistent local storage
	db bolt.DB
}

func (s *PersistentStorage) Set(key, value string) {
	defer func() {
		e := recover()
		fmt.Printf("Set error on key,value: %s, %s, error:%s ", key, value, e)
	}()

	tx, err := s.db.Begin(true)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	// covert []byte of key,value and PUT in db
	keyBytes := []byte(key)
	valBytes := []byte(value)
	bucket := tx.Bucket(gokv_store)
	if err := bucket.Put(keyBytes, valBytes); err != nil {
		panic(err)
	}

	e := tx.Commit()
	if e != nil {
		panic(e)
	}
}

func (s *PersistentStorage) Get(key string) string {
	defer func() {
		e := recover()
		fmt.Printf("Get error on key:%s, err:%s ", key, e)
	}()

	tx, err := s.db.Begin(false)
	if err != nil {
		panic(err)
	}

	bucket := tx.Bucket(gokv_store)
	valueByte := bucket.Get([]byte(key))
	return string(valueByte)
}

func (s *PersistentStorage) Exist(key string) bool {
	defer func() {
		e := recover()
		fmt.Printf("Exist error on key:%s, err:%s ", key, e)
	}()

	tx, err := s.db.Begin(false)
	if err != nil {
		panic(err)
	}

	bucket := tx.Bucket(gokv_store)
	valueByte := bucket.Get([]byte(key))
	return len(valueByte) != 0
}

func (s *PersistentStorage) Remove(key string) bool {
	defer func() {
		e := recover()
		fmt.Printf("Exist error on key:%s, err:%s ", key, e)
	}()

	tx, err := s.db.Begin(false)
	if err != nil {
		panic(err)
	}

	bucket := tx.Bucket(gokv_store)
	if err := bucket.Delete([]byte(key)); err != nil {
		panic(err)
	}

	return true
}
