package main

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFs embed.FS

func copyFilefromTemplate(templatePath, targetFile string) error {

	if fileExists(targetFile) {
		return errors.New(targetFile + " already exists!")
	}

	// 讀取同層資料夾下的資料
	data, err := templateFs.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	// 複製讀取到的檔案到傳進來的路徑
	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {

	// 複製檔案到 to 這個路徑 最後一個參數為權限
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}

	return true
}
