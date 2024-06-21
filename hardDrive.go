package main

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func allHardDrive(w2 fyne.Window) *fyne.Container {
	leftDriver := container.NewVBox()
	lsblkArr := []string{}
	cmdLs := exec.Command("lsblk", "-o", "NAME,MOUNTPOINT")
	var out strings.Builder
	cmdLs.Stdout = &out
	err := cmdLs.Run()
	if err != nil {
		log.Fatal(err)
	}
	splitedSLsbk := strings.Split(out.String(), "\n")
	lsblkArr = append(lsblkArr, splitedSLsbk...)
	for _, lsblk := range lsblkArr {
		name, path, _ := strings.Cut(lsblk, " ")
		if strings.ContainsAny(name, "loop") || strings.ContainsAny(name, "NAME") || !strings.ContainsAny(path, "/") || path == "/boot/efi" {
			continue
		}
		icon := getIcon(IconStruct{source: "./images/drive_icon.png"})
		button := widget.NewButtonWithIcon(name, icon, func() {
			strp := &str
			*strp = path
			go func() {
				time.Sleep(time.Millisecond)
				openFileFunc(w2 , a)
			}()
		})
		leftDriver.Add(button)
	}
	return leftDriver
}