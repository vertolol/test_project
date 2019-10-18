package main

import "log"

func ifPanic(err error) {
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
}
