package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, cruel world!")
	go updateMessage("Hello, cruel world2!")
	wg.Wait()

	if msg != "Hello, cruel world!" {
		t.Error("incorrect value in msg")
	}
}
