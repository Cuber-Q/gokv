package strategy

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_KeySharding(t *testing.T) {
	s_1 := "a:fadf:q"
	s_2 := "cc:erqe:da"
	s_3 := "cc:erqe:da"

	st := &HashShardStrategy{}
	h_1, _ := st.KeySharding(s_1)
	h_2, _ := st.KeySharding(s_2)
	h_3, _ := st.KeySharding(s_3)

	fmt.Println(h_1)
	fmt.Println(h_2)
	fmt.Println(h_3)

	assert.True(t, h_1 > 0, "h_1 > 0")
	assert.True(t, h_2 > 0, "h_2 > 0")
	assert.True(t, h_3 > 0, "h_3 > 0")
	assert.True(t, h_1 != h_2, "h_1 != h_2")
	assert.True(t, h_2 == h_3, "h_2 == h_3")
}

func Test_channel_multi_read(t *testing.T) {
	ch := make(chan string)
	go func() {
		for {
			x := <-ch
			fmt.Println("func_1: ", x)
		}
	}()
	go func() {
		for {
			x := <-ch
			fmt.Println("func_2: ", x)
		}
	}()
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
	ch <- "7"
	ch <- "8"
	ch <- "9"
	ch <- "10"
}

func Test_for_chaos(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second)
}

func Test_for_all_good(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second)
}

func Test_for_range(t *testing.T) {
	//arr := []int{} // init as empty arr, can't be assign anymore
	var arr []int
	arr = make([]int, 10)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	for k, v := range arr {
		fmt.Println(k, ":", v)
	}

	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"
	m["k4"] = "v4"

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func Test_slice_copy(t *testing.T) {
	s := make([]int, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	fmt.Println(s[2:])
	s_copy := s[2:]
	fmt.Println(s_copy)
}

func Test_bufwrite(t *testing.T) {
	file, e := os.Open("/test.txt")
	if e != nil {
		fmt.Errorf("%s", e)
		return
	}

	writer := bufio.NewWriter(file)
	writer.Write([]byte{})
	writer.Flush()
}

func Test_defer(t *testing.T) {
	result := 0
	defer func() {
		result++
	}()
	return
}

func Test_panic(t *testing.T) {
	defer func() {
		e := recover()
		fmt.Println(e)
	}()

	panic("error")
}
