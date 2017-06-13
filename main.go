package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("pbpaste").Output()
	s := strings.Split(string(out), ",")
	if err != nil {
		fmt.Print(err.Error())
	}

	for _, v := range s {
		url := fmt.Sprintf("https://senseye.atlassian.net/browse/SIGHT-%s", strings.TrimSpace(v))
		err = exec.Command("open", url).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%+v", s)
}
