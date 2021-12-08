package util

import "fmt"

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
