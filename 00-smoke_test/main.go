// Copyright (c) 2023 Luka Ivanović
// This code is licensed under The MIT Licence (see LICENCE for details).

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mirror(connection net.Conn) {
	defer connection.Close()
	if _, err := io.Copy(connection, connection); err != nil {
		log.Println("Error copying:", err.Error())
		return
	}
}

func getEnvOrDefault(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return value
}

func main() {
	ip := getEnvOrDefault("IP", "127.0.0.1")
	port := getEnvOrDefault("PORT", "8080")
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatalln("Error listening:", err)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s:%s\n", ip, port)
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go mirror(connection)
	}
}
