package main

import (
	"fmt"
	"os"

	"github.com/wasuken/scout/apt"
	"github.com/wasuken/scout/config"
	"github.com/wasuken/scout/pacman"
	"github.com/wasuken/scout/send"
)

func main() {
	pacMan := os.Args[1]
	err, config := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	if pacMan == "pacman" {
		err, info := pacman.GetInfo()
		if err != nil {
			panic(err)
		}
		send.SendSrv(info, config.URL)
	} else if pacMan == "apt" {
		apt.GetInfo()
		err, info := apt.GetInfo()
		if err != nil {
			panic(err)
		}
		send.SendSrv(info, config.URL)
	} else {
		fmt.Println("Ha?")
	}
}
