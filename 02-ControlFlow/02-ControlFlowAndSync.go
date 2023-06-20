package main

import (
	"fmt"
	"time"
)

type Server struct {
	// Empty struct will be passed to tell code to break the server loop.
	quitCh  chan struct{}
	msgChan chan string
}

func newServer() *Server {
	return &Server{
		quitCh:  make(chan struct{}),
		msgChan: make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("Server Started")
	s.loop()
}

func (s *Server) sendMessage(msg string) {
	s.msgChan <- msg
}

func (s *Server) quitServer() {
	s.quitCh <- struct{}{}
	fmt.Println("Server Shutdown")
}

func (s *Server) loop() {
mainLoop:
	for {
		select {
		case <-s.quitCh:
			break mainLoop
		case msg := <-s.msgChan:
			s.handleMessage(msg)
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Message: ", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quitServer()
	}()

	server.start()
}
