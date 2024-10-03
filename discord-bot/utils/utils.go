package utils

import "log"

func CheckNilErr(e error) {
	if e != nil {
	 log.Fatal("Error message")
	}
 }