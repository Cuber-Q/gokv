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

func ValidAddress(address string) bool {
	add := strings.Split(address,":")
	if len(add) != 2 {
		return false
	}

	if !ValidIp(add[0]) {
		return false
	}

	if _, e := strconv.Atoi(add[1]);e != nil {
		return false
	}

	i, _ := strconv.Atoi(add[1])
	if i <= 0 || i > 65535 {
		return false
	}

	return true
}