package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	if len(os.Args) >= 2 {
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		repl.StartFile(bytes, os.Stdout)
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}