package main

import "os"

const parsePath = "C:\\Users\\AdvikB\\Downloads\\WhatsApp Chat with Nigga"

func main() {
	err := os.Link(parsePath, "Whats")
	go func() {
		if err != nil {
			panic(err)
		}
	}()
	recover()
	file, err := os.OpenFile(parsePath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()

}
