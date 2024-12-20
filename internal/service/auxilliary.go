package service

import (
	"log"
)

func FailOnErrors(err error, message string) {
	if err != nil {
		log.Printf("%s: %s", err, message)
	}
}
