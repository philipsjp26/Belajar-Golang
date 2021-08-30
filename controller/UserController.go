package controller

import "fmt"

func Hello(name string) string {
	message:= fmt.Sprintf("Hai %v tempoek ", name)
	return message
}