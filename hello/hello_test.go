package main

import "testing"

func TestReceive(t *testing.T) {
	receive()
}

func TestSend(t *testing.T) {
	send()
}

func TestLog(t *testing.T) {
	go func() {
		test()
	}()
}
