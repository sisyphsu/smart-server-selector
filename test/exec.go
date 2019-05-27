package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	println("> ssh -p876 sull@12.22.33.44")
	cmd := exec.Command("ssh", "13.127.159.17")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("error: ", err)
	}
}
