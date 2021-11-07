package cpu

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/wasuken/scout/send"
)

func getCPUInfo() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val
				if i == 4 {
					idle = val
				}
			}
			return
		}
	}
	return
}

func calcCPUPercent() float64 {
	idle0, total0 := getCPUInfo()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUInfo()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	return 100 * (totalTicks - idleTicks) / totalTicks
}

func GetInfo() (error, send.SendCPUInfo) {
	time := calcCPUPercent()
	name, err := os.Hostname()
	if err != nil {
		return err, send.SendCPUInfo{}
	}

	return nil, send.SendCPUInfo{
		Name:        name,
		Arch:        runtime.GOARCH,
		PackManType: "cpu",
		CPUTime:     time,
	}
}
