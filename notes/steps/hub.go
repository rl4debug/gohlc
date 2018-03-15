package main

import (
	"bufio"
	"fmt"
	"net"
)

//Listen at port
//Client connect to Hub -> Save all information
//Broadcast

type Hub struct {
	clients []Client
}

type Client struct {
}

type Server struct {
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var c = make(chan int)
	close(c)
	if c == nil {
		fmt.Println("Nil roi")
	}

	fmt.Println("hehe")

	return

	ln, err := net.Listen("tcp", ":9989")
	if err != nil {
		// handle error
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			panic(err)
		}

		// fmt.Println("Connected")
		// _, err = conn.Write([]byte(fmt.Sprint("Accepted client: ", conn.RemoteAddr().String(), '\n')))

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("Sending accepting message.")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		msg := fmt.Sprint("Received: ", text)
		fmt.Println(msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			panic(err)
		}
	}
}
