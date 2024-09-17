package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var idList = make(map[int]time.Time)

func contains(list map[int]time.Time, id int) bool {
	if _, ok := list[id]; ok {
		// if the time difference is more than 15 minutes, return true
		if time.Since(list[id]).Minutes() > 15 {
			delete(idList, id)
			return false
		}
		return true
	}
	return false
}

func GenerateId() (rnd6Digits string, err error) {
	// 	generate a random number between 0 and 999999
	rnd := rand.Intn(999999)

	// parse 0s to the left
	rnd6Digits = fmt.Sprintf("%06d", rnd)

	if err != nil {
		return
	}

	// 	check if the number is in the idList
	if contains(idList, rnd) {
		// 	if it is, generate a new number
		rnd6Digits, err = GenerateId()

		var idInt int = 0
		idInt, err = strconv.Atoi(rnd6Digits)

		if err != nil {
			return
		}

		delete(idList, idInt)
		return
	} else {
		// 	if it isn't, add it to the idList
		idList[rnd] = time.Now()
		return
	}
}

func CheckifIdIsExpired(id int) bool {
	return contains(idList, id)
}
