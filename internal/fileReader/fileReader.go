package filereader

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Filereader(a fyne.App) {
	w := a.NewWindow("file manager")
	file_Dialog := dialog.NewFileOpen(
		func(r fyne.URIReadCloser, _ error) {
			data, _ := io.ReadAll(r)
			result := fyne.NewStaticResource("name", data)
			entry := widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))	
			w := fyne.CurrentApp().NewWindow(
				string(result.StaticName))
			w.SetContent(container.NewScroll(entry))
		}, w)
	fmt.Println(file_Dialog)

	w.Resize(fyne.NewSize(400, 400))
	w.Show()
	file_Dialog.Show()
}

