package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sysconf/pacman"
)

func main() {
	fmt.Println("updating system configuration...")

	var packages []pacman.Package

	pfile, err := ioutil.ReadFile("packages.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(pfile, &packages)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !pacman.Exists() {
		os.Exit(1)
	}

	fmt.Printf("packages to include %v\n", packages)

	installedPackages := pacman.InstalledPackages()

	fmt.Printf("number of installed packages: %v\n", len(installedPackages))

	pacman.Update(packages)

	fmt.Println("deps for firefox:", pacman.Package{"firefox"}.Dependencies())
	fmt.Println("deps for gimp:", pacman.Package{"gimp"}.Dependencies())

}
