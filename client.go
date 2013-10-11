/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:27 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
    "fmt"
    "bufio"
    "net"
)

type Client struct {
	incoming chan string
	outgoing chan string // only for maze data in json format
	echoer chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
	    fmt.Println("Reading..")
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
		client.writer.WriteString("\n\n----------------------------------------\n\n")
		client.writer.Flush()
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString("JSON: \n" + data + "\n\n----------------------------------------\n\n")
		client.writer.Flush()
	}
}

func (client *Client) Echo() {
	for data := range client.echoer {
		client.writer.WriteString("ECHO: \n" + data + "\n\n----------------------------------------\n\n")
		client.writer.Flush()
	}
}


func (client *Client) Listen() {
	go client.Read()
	go client.Write()
	go client.Echo()
}

func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		echoer: make(chan string),
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}

