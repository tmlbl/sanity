package main

import (
	"path/filepath"
)

// Config represents a Sanity blessed config file for a package
type Config struct {
	path   string // The local path from the installation directory
	prefix string // The top-level path under which it will be placed
}

// Package represents a system package and its built-in configuration
type Package struct {
	Installs []string // Apt package names to be installed
	Configs  []Config // Paths to configs to be added
}

// Install performs all installation steps for the Package
func (p *Package) Install() {
	aptInstall(p.Installs...)
	for _, c := range p.Configs {
		err := copyOver(c.path, filepath.Join(c.prefix, c.path), false)
		if err != nil {
			panic(err)
		}
	}
}

var lightDM = Package{
	Installs: []string{"lightdm", "lightdm-gtk-greeter"},
	Configs: []Config{
		Config{
			path:   "etc/lightdm/lightdm-gtk-greeter.conf",
			prefix: "/",
		},
	},
}

var openbox = Package{
	Installs: []string{"openbox"},
	Configs: []Config{
		Config{
			path:   "etc/xdg/openbox/rc.xml",
			prefix: "/",
		},
		Config{
			path:   "etc/xdg/openbox/autostart",
			prefix: "/",
		},
	},
}

var vim = Package{
	Installs: []string{"vim"},
	Configs: []Config{
		Config{
			path:   "etc/vim/vimrc.local",
			prefix: "/",
		},
	},
}

var lilyterm = Package{
	Installs: []string{"lilyterm"},
	Configs: []Config{
		Config{
			path:   "etc/xdg/lilyterm.conf",
			prefix: "/",
		},
	},
}
