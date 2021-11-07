package main

import (
	"os"

	"github.com/wasuken/scout/apt"
	"github.com/wasuken/scout/config"
	"github.com/wasuken/scout/cpu"
	"github.com/wasuken/scout/pacman"
)

func main() {
	main_cmd := os.Args[1]
	err, config := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	switch main_cmd {
	case "pacman":
		err, info := pacman.GetInfo()
		if err != nil {
			panic(err)
		}
		info.SendSrv(info, config.URL)
	case "apt":
		err, info := apt.GetInfo()
		if err != nil {
			panic(err)
		}
		info.SendSrv(info, config.URL)
	case "cpu":
		err, info := cpu.GetInfo()
		if err != nil {
			panic(err)
		}
		info.SendSrv(info, config.URL)
	default:
		println("ha?")
	}
}
