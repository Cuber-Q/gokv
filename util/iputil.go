package util

import (
	"strings"
	"log"
	"strconv"
)

func ValidIp(ip string) bool {
	if ! strings.Contains(ip, ".") {
		log.Fatal("invalid ip value: ", ip)
		return false
	}

	numbers := strings.Split(ip, ".")

	for i := 0; i < 4; i++ {
		if i == 0 {
			if len(numbers[i]) == 0 {
				log.Println("invalid ip value: ", ip)
				return false
			}

			n, e := strconv.Atoi(numbers[i])
			if e != nil {
				log.Println("invalid ip value: ", e)
				return false
			}

			if n <= 0 || n > 255 {
				log.Println("invalid ip value: ", ip)
				return false
			}

		}

		n, e := strconv.Atoi(numbers[i])
		if e != nil {
			log.Println("invalid ip value: ", e)
			return false
		}

		if n < 0 || n > 255 {
			log.Println("invalid ip value: ", ip)
			return false
		}
	}
	return true
}
