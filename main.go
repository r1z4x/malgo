package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	plugins := LoadPlugins("./plugins")
	fmt.Println(plugins)
	for _, plugin := range plugins {
		plugin.Run("google.com")
	}
}

type MyPlugin interface {
	Run(host string)
	Register()
}

func LoadPlugins(path string) []MyPlugin {
	/*var plugins []Plugin
	files, _ := filepath.Glob(path + "/*.so")
	for _, file := range files {
		p, err := plugin.Open(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		symPlugin, err := p.Lookup("Plugin")
		if err != nil {
			fmt.Println(err)
			continue
		}
		var plugin Plugin
		plugin, ok := symPlugin.(Plugin)
		if !ok {
			fmt.Println("Unexpected type from module symbol")
			continue
		}
		plugins = append(plugins, plugin)
	}
	return plugins*/

	fmt.Println("Loading plugins...")

	var plugins []MyPlugin

	root, _ := os.Getwd()
	pluginsDir := filepath.Join(root, "plugins")

	err := filepath.Walk(pluginsDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".so" {
			p, err := plugin.Open(path)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			registerSym, err := p.Lookup("Register")
			fmt.Println(registerSym)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			// Register fonksiyonunu çağırın

			register, ok := registerSym.(func() MyPlugin)
			if !ok {
				fmt.Println("Invalid plugin")
				return nil
			}

			plugin := register()
			plugins = append(plugins, plugin)

		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return plugins
}
