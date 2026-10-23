package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtInCommands = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

func main() {

	for {
		fmt.Print("$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		if cmd == "exit\n" {
			break
		}
		if strings.HasPrefix(cmd, "echo ") {
			fmt.Println(cmd[5 : len(cmd)-1])
			continue
		}
		if strings.HasPrefix(cmd, "type ") {
			if _, ok := builtInCommands[cmd[5:len(cmd)-1]]; ok {
				fmt.Printf("%s is a shell builtin\n", cmd[5:len(cmd)-1])
			} else if path, err := exec.LookPath(cmd[5 : len(cmd)-1]); err == nil {
				fmt.Printf("%s is %s\n", cmd[5:len(cmd)-1], path)
			} else {
				fmt.Printf("%s: not found\n", cmd[5:len(cmd)-1])
			}
			continue
		}
		fmt.Printf("%s: command not found\n", cmd[:len(cmd)-1])
	}
}
