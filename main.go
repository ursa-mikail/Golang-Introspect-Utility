package main

import (
	"fmt"

	"example.com/demo/lib/code_inspect"
)

func main() {
	filePath := "./data/example.go"
	jsonComment, err := code_inspect.ExtractJSONComment(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	code_inspect.MainMenu(jsonComment)
}

/*
% go mod init example.com/demo
% go run main.go

*/
