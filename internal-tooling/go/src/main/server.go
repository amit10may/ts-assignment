package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	name string
	conn net.Conn
}

var requests []string 

func sendStorageServerResponse(clientJobs chan ClientJob) {
	for {
		// Wait for the next job to come off the queue.
		clientJob := <-clientJobs

		requests = append(requests, clientJob.name)
		fmt.Println("Request Count: ", len(requests));

		// Send back the response.
		clientJob.conn.Write([]byte("Saved " + clientJob.name ))
	}
}

func sendListingServerResponse(clientJobs chan ClientJob) {
	for {
		// Wait for the next job to come off the queue.
		clientJob := <-clientJobs

		fmt.Println("Received command: ", clientJob.name)
		
		// Send back the response.
		for _,request := range requests {
			clientJob.conn.Write([]byte(request))
		}
	}
}

func main() {
   go startStorageServer()
   startListingServer()
}

func startStorageServer() {
	clientJobs := make(chan ClientJob)
	go sendStorageServerResponse(clientJobs)
	
	ln, err := net.Listen("tcp", ":8080")
	check(err, "Storage server is ready at 8080.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}


func startListingServer() {
	clientJobs := make(chan ClientJob)
	go sendListingServerResponse(clientJobs)
	
	ln, err := net.Listen("tcp", ":8081")
	check(err, "Listing server is ready at 8081.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}
