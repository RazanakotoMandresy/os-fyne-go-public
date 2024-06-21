package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func buttonSpace(lsArrRanged string) *widget.Button {
	buttonForSpace := widget.NewButton("", func() {
		cmdLs := exec.Command("ls", "-l", str+"/"+lsArrRanged)
		var out strings.Builder
		cmdLs.Stdout = &out
		err := cmdLs.Run()
		if err != nil {
			log.Fatal(err)
		}
		content := out.String()
		theContP := &theCont
		*theContP = content
		info := spaceShow(content)
		fmt.Println(info)
	})
	return buttonForSpace
}
func spaceShow(cont string) string {
	return cont
}
func infoSpace() *widget.Card {
	images := canvas.NewImageFromFile("./images/folderImages.jpg")
	cont := spaceShow(theCont)
	images.FillMode = canvas.ImageFillOriginal
	infoStr := widget.NewLabel(cont)
	info := widget.NewCard("", "information", container.New(layout.NewVBoxLayout(), images, infoStr))
	return info
}
