package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

func main() {
	port := rand.Int31n(65535)
	listen, err := net.Listen("tcp", "127.0.0.1:" + strconv.Itoa(int(port)))
	if err != nil {
		fmt.Println("listen failed, err " + err.Error())
	}
	fmt.Println("listen port " + strconv.Itoa(int(port)))
	for {
		conn, _ := listen.Accept()
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from conn failed, err ", err)
			break
		}

		recv := string(buf[:n])
		fmt.Printf("receive %v\n", recv)

		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Println("write from conn failed, err ", err)
			break
		}
	}
}
