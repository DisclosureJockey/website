package main

import (
	"crypto/rand"
	"fmt"
)

var guids = make(chan string, 5)

func init() {
	go func() {
		for {
			b := make([]byte, 16)
			_, err := rand.Read(b)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			guids <- fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
		}
	}()
}
