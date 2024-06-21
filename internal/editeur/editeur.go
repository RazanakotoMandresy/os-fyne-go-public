package editeur

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var sContent string

type IconStruct struct {
	source string
}

func getIcon(I IconStruct) *fyne.StaticResource {
	bytes, _ := os.ReadFile(I.source)
	stringByte := string(bytes)
	theIcon := fyne.NewStaticResource("", []byte(stringByte))
	return theIcon
}

func Editeur(a fyne.App, val string) {

	w5 := a.NewWindow("word")
	var bools bool

	editor := widget.NewMultiLineEntry()
	editor.Text = val
	italicIcon := getIcon(IconStruct{source: "./images/italic.png"})
	boldIcon := getIcon(IconStruct{source: "./images/boldIcon.png"})
	monspaceIcon := getIcon(IconStruct{"./images/Monospace_font.svg"})
	saveIcon := getIcon(IconStruct{source: "./images/save.png"})
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(boldIcon, func() {
			editor.TextStyle.Bold = !bools
		}),
		widget.NewToolbarAction(italicIcon, func() {
			editor.TextStyle.Italic = !bools
		}),
		widget.NewToolbarAction(monspaceIcon, func() {
			editor.TextStyle.Monospace = !bools
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(saveIcon, func() {
			surEnregister(sContent, a)
		}),
	)
	editor.OnChanged = func(s string) {
		textValue(s)
	}

	content := container.NewBorder(toolbar, nil, nil, nil, editor)
	w5.SetContent(content)
	w5.Resize(fyne.NewSize(800, 800))
	w5.Show()
}

func surEnregister(stringContent string, a fyne.App) {
	w6 := a.NewWindow("sur d'enregister")
	fileName := widget.NewEntry()
	fileName.PlaceHolder = "Le titre de votre document"
	selectExt := widget.NewSelect([]string{".txt", ".docx", ".odt", ".doc"}, func(valeurExt string) {
		fileName.OnSubmitted = func(s string) {
			fileNameCree, err := os.Create(s + valeurExt)
			if err != nil {
				fmt.Println(err)
			}
			fileNameCree.WriteString(stringContent)
		}
	})
	content := container.NewHBox(fileName, selectExt)
	w6.SetContent(content)
	w6.Resize(fyne.NewSize(400, 400))
	w6.Show()
}
func textValue(s string) string {
	sContentP := &sContent
	*sContentP = s
	return s
}
