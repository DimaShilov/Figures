package service

import "fmt"

func WrapError(desc string, err error) {
	if err != nil {
		fmt.Println(desc, err)
	}
}
