package core

import (
	"log"
	"sync"
)

// core storage

var store = make(map[string]string)
var mutex sync.Mutex

func Set(key, value string) string {
	mutex.Lock()
	defer mutex.Unlock()

	if key == "" {
		log.Println("Set \t[IGNORED], key=", key, ", value=", value)
		return "ERROR: <key> is empty"
	}
	store[key] = value
	log.Println("Set \t[SCUESS], key=", key, ", value=", value)
	return "OK"
}

func Get(key string) string {
	mutex.Lock()
	defer mutex.Unlock()

	value := store[key]
	log.Println("Get \t[SCUESS], key=", key, ", value=", value)
	return value
}

func Exist(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	exist := store[key] != ""
	log.Println("Exist \t[SCUESS], key=", key, ", exist=", exist)
	return exist
}

func Keys() []string {
	mutex.Lock()
	defer mutex.Unlock()

	keys := make([]string, len(store))
	for key := range store {
		if key != "" {
			keys = append(keys, key)
		}
	}
	log.Println("Keys \t[SCUESS], keys=", keys)
	return keys
}
