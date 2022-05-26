package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFs embed.FS

func copyFilefromTemplate(templatePath, targetFile string) error {
	// TODO

	data, err := templateFs.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
