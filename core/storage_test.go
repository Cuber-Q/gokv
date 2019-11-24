package core

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNewStorage(t *testing.T) {
	s := NewStorage()
	count := make(chan int64, 1e6)

	go func() {
		for i := 0; i < 1e6; i++ {
			cmd := &StoreCMD{
				T:      SET,
				K:      strconv.Itoa(i),
				V:      strconv.Itoa(i),
				RespCh: make(chan string, 1),
			}
			start := time.Now().UnixNano()
			s.SetV3(cmd)

			<-cmd.RespCh
			cost := time.Now().UnixNano() - start
			//log.Printf("SET invoked: %v, cost: %vns", cmd.k, cost)
			count <- cost
		}
	}()

	go func() {
		var idx = int64(0)
		var t = int64(0)
		for {
			i := <-count
			t += i
			idx++
			if idx >= 1e6 {
				avg := t / idx
				log.Printf("avg cost time: %vns, t=%v, idx=%v", avg, t, idx)
				break
			}
		}

	}()

	time.Sleep(1000 * 1000 * time.Second)
}
