package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func displayImgages(a fyne.App, str string) {
	w3 := a.NewWindow(str)
	images := canvas.NewImageFromFile(str)
	images.FillMode = canvas.ImageFillStretch
	w3.SetContent(images)
	w3.Resize(fyne.NewSize(400, 400))
	w3.Show()
}
