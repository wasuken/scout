package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PackageInfo struct {
	Name        string
	Version     string
	Description string
}

type SendPackageInfo struct {
	Packs       []PackageInfo
	Name        string // サーバ名
	PackManType string // パッケージマネージャの種類(apt|pacman)
	Arch        string // サーバのOSのArch
}

type SendCPUInfo struct {
	CPUTime     float64
	Name        string // サーバ名
	PackManType string // パッケージマネージャの種類(apt|pacman)
	Arch        string // サーバのOSのArch
}

func (cpuinfo SendCPUInfo) SendSrv(info SendCPUInfo, url string) {
	json, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(json)
	coreSendSrv(b, url)
}

func (packinfo SendPackageInfo) SendSrv(info SendPackageInfo, url string) {
	json, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(json)
	coreSendSrv(b, url)
}
func coreSendSrv(jsonb *bytes.Buffer, url string) {
	req, err := http.NewRequest(
		"POST",
		url,
		jsonb,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var buf []byte
	_, err = resp.Body.Read(buf)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
	fmt.Println("finished.")
}
