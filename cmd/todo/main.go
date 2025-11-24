package main

import (
	//"fmt"
	"os"
	//"strconv"

	"example.com/todo-cli/internal/todo"
)

func main(){
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	store:=todo.NewStore("data/todos.json")
	cmd:=os.Args[1]

	switch cmd {
		case "add":
			text:=os.Args[2]
			err:=store.Add(text)
			if(err!=nil){
				os.Exit(1)
			}
	}
}
