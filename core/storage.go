package core

import (
	"encoding/json"
	"io"
	"log"
	"sync"
)

type StorageCmdType uint

const (
	SET    StorageCmdType = 0
	GET    StorageCmdType = 1
	EXIST  StorageCmdType = 2
	REMOVE StorageCmdType = 3
)

type Storage struct {
	store map[string]string
	mutex sync.Mutex
	sync.RWMutex
	cmdCh chan *StoreCMD
}

type StoreCMD struct {
	T      StorageCmdType
	K      string
	V      string
	RespCh chan string
}

func NewStorage() *Storage {
	s := &Storage{
		store: make(map[string]string),
		cmdCh: make(chan *StoreCMD, 1000*1000),
	}
	go s.runCmd()
	return s
}

func (s *Storage) SetV2(key, value string) {
	//mutex.Lock()
	s.store[key] = value
	//mutex.Unlock()
	log.Printf("storage SET OK. k=%v, v=%v", key, value)
}

func (s *Storage) GetV2(key string) string {
	value := s.store[key]
	log.Printf("storage GET OK. k=%v, v=%v", key, value)
	return value
}

// Marshal serializes cache data
func (s *Storage) Marshal() ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	dataBytes, err := json.Marshal(s.store)
	return dataBytes, err
}

// UnMarshal deserializes cache data
func (s *Storage) UnMarshal(serialized io.ReadCloser) error {
	var newData map[string]string
	if err := json.NewDecoder(serialized).Decode(&newData); err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	s.store = newData

	return nil
}

func (s *Storage) SetV3(cmd *StoreCMD) {
	s.cmdCh <- cmd
}

// async exec concurrent request and notify invoker when completed
func (s *Storage) runCmd() {
	for {
		cmd := <-s.cmdCh
		t := cmd.T
		if t == SET {
			s.SetV2(cmd.K, cmd.V)
			cmd.RespCh <- "ok"
		} else if t == GET {
			cmd.RespCh <- s.GetV2(cmd.K)
		}
	}

}
