package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type Config struct {
	UI UIOpts `yaml:"ui"`
	Menu []Entry `yaml:"menu"`
	Binds Binds `yaml:"binds"`
}

type UIOpts struct {
	ShowPath bool `yaml:"show_path"`
	ShowLines bool `yaml:"show_lines"`
	ShowHints bool `yaml:"show_hints"`
	SubmenuSymbol string `yaml:"submenu_symbol"`
}

type Entry struct {
	Key string `yaml:"key"`
	Desc string `yaml:"desc"`
	Cmd string `yaml:"cmd,omitempty"`
	Submenu []Entry `yaml:"submenu,omitempty"`
}

type Binds struct {
	Quit string `yaml:"quit"`
	Back string `yaml:"back"`
}

var ui UIOpts
var binds Binds
var oldState *term.State
var fd int

func main() {
	cfg := loadConfig()
	ui = cfg.UI
	binds = cfg.Binds

	if ui.SubmenuSymbol == "" { ui.SubmenuSymbol = "+" }
	if binds.Quit == "" { binds.Quit = "q" }
	if binds.Back == "" { binds.Back = "h" }

	fd = int(os.Stdin.Fd())
	var err error
	oldState, err = term.MakeRaw(fd)
	if err != nil { panic(err) }
	defer fix() // oaoaooaoaoaoaoa

	_ = runMenu(cfg.Menu, "/")
	clear()
}


func fix() { if oldState != nil { term.Restore(fd, oldState) } } // the bloody thing wont do the thing aaaaaaaaaa

func loadConfig() Config {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config/keyselect/config.yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Could not read config:", err)
		os.Exit(1)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Println("Bad YAML:", err)
		os.Exit(1)
	}

	return cfg
}

func runMenu(menu []Entry, path string) bool {
	for {
		drawMenu(menu, path)

		key := readKey()

		if key == binds.Quit {
			clear()
			return true // true just quits
		}

		if key == binds.Back { return false } // false raises a layer

		for _, e := range menu {
			if e.Key != key { continue }

			if len(e.Submenu) > 0 {
				newPath := path + e.Desc + "/"
				if runMenu(e.Submenu, newPath) { return true }
				break
			}

			if e.Cmd != "" {
				clear()
				runCmd(e.Cmd)
				return false
			}
		}
	}
}


func drawMenu(menu []Entry, path string) {
	clear()

	if ui.ShowPath { fmt.Printf("➤ %s\r\n", path) }

	if ui.ShowLines {
		fmt.Println(strings.Repeat("─", 20))
		fmt.Printf("\r") // bodge primus
	}

	for _, e := range menu {
		if len(e.Submenu) > 0 {
			fmt.Printf("[%s] %s %s\r\n", e.Key, ui.SubmenuSymbol, e.Desc)
		} else {
			fmt.Printf("[%s]   %s\r\n", e.Key, e.Desc)
		}
	}

	if ui.ShowLines {
		fmt.Println(strings.Repeat("─", 20))
		fmt.Printf("\r") 
	}

	if ui.ShowHints { fmt.Printf("%s = back    %s = quit\r\n", binds.Back, binds.Quit) }
}

func runCmd(cmd string) {
	c := exec.Command("sh", "-c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil { fmt.Println("Command failed:", err) }
}

func readKey() string {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	if err != nil { return "" }
	return string(buf)
}

func clear() { fmt.Print("\033[H\033[2J") }
