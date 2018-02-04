package main

import "fmt"

var corePackages = []string{
	"sudo",
	"rsync",

	"xorg",

	"git",

	"chromium",

	"redshift",
	"nitrogen",
	"htop",
	"build-essential",
}

func main() {
	fmt.Println("Beginning the installation.")
	runCmd("apt-get", "update")

	// Install core packages
	aptInstall(corePackages...)

	openbox.Install()
	lightDM.Install()
	vim.Install()
	lilyterm.Install()

	// Finished. Ask the user if they'd like to reboot
	if askYesNo("Would you like to reboot the machine now?", true) {
		runCmd("shutdown", "-r", "now")
	}
}
