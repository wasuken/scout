package apt

import (
	"os"

	"runtime"

	"github.com/arduino/go-apt-client"
	"github.com/wasuken/scout/send"
)

func GetInfo() (error, send.SendPackageInfo) {
	apt.CheckForUpdates()
	allPkgs, err := apt.List()
	if err != nil {
		panic(err)
	}
	pkgInfos := []send.PackageInfo{}
	for _, pkg := range allPkgs {
		pack := send.PackageInfo{Name: pkg.Name, Version: pkg.Version}
		pkgInfos = append(pkgInfos, pack)
	}
	name, err := os.Hostname()
	if err != nil {
		return err, send.SendPackageInfo{}
	}
	return nil, send.SendPackageInfo{
		Name:        name,
		PackManType: "apt",
		Arch:        runtime.GOARCH,
		Packs:       pkgInfos,
	}
}
