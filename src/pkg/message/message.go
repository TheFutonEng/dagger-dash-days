package message

import "fmt"

func TellUserOnFail(errorStr string, action string) {
	fmt.Println("Error: " + errorStr)
	fmt.Println("Recommendation: " + action)
}

func Info(info string) {
	fmt.Println(info)
}