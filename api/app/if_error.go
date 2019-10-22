package app

import "log"

func IfError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
