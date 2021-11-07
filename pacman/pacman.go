package pacman

import (
	"os"
	"runtime"

	"github.com/goulash/pacman"
	"github.com/goulash/pacman/pkgutil"
	"github.com/wasuken/scout/send"
)

func GetInfo() (error, send.SendPackageInfo) {
	localPkgs, err := pacman.ReadLocalDatabase(func(er error) error {
		panic(er)
	})
	if err != nil {
		return err, send.SendPackageInfo{}
	}
	localPkgMap := pkgutil.MapPkg(localPkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})

	name, err := os.Hostname()
	if err != nil {
		return err, send.SendPackageInfo{}
	}

	pkgInfos := []send.PackageInfo{}
	for _, pkg := range localPkgMap {
		pkgInfo := send.PackageInfo{
			Name:        pkg.PkgName(),
			Version:     pkg.Version,
			Description: pkg.Description}
		pkgInfos = append(pkgInfos, pkgInfo)
	}
	return nil, send.SendPackageInfo{
		Name:        name,
		Arch:        runtime.GOARCH,
		PackManType: "pacman",
		Packs:       pkgInfos,
	}
}
