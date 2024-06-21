package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/RazanakotoMandresy/os-fyne-go/internal/editeur"
	filereader "github.com/RazanakotoMandresy/os-fyne-go/internal/fileReader"
)

type IconStruct struct {
	source string
}
type buttonCmd struct {
	name        string
	ic          *fyne.StaticResource
	commandName string
}

func execCommande(b buttonCmd) *widget.Button {
	btn := widget.NewButtonWithIcon(b.name, b.ic, func() {
		cmd := exec.Command(b.commandName)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})
	return btn
}
func getIcon(I IconStruct) *fyne.StaticResource {
	bytes, _ := os.ReadFile(I.source)
	stringByte := string(bytes)
	theIcon := fyne.NewStaticResource("", []byte(stringByte))
	return theIcon
}
func updateTime(clock, date *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	dates := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	clock.SetText(formatted)
	date.SetText(dates.String())
}

var a = app.New()

func main() {
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	date := widget.NewLabel("")
	updateTime(clock, date)
	top := container.New(layout.NewVBoxLayout(), date, clock)
	onTop := canvas.NewPositionAnimation(fyne.NewPos(900, 0), fyne.NewPos(900, 0), time.Second, top.Move)
	onTop.Start()
	dashboard := dashboard(a)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock, date)
		}
	}()
	content := container.New(layout.NewVBoxLayout(), top, dashboard)
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
func dashboard(a fyne.App) *fyne.Container {
	codeIcons := getIcon(IconStruct{source: "./images/logo.png"})
	openCodeBtn := execCommande(buttonCmd{name: "vsCode", ic: codeIcons, commandName: "code"})
	openBrave := getIcon(IconStruct{source: "./images/brave.png"})
	openBraveBtn := execCommande(buttonCmd{name: "brave", ic: openBrave, commandName: "brave"})
	openFiles := getIcon(IconStruct{source: "./images/folder_icon.png"})
	openFilesBtn := widget.NewButtonWithIcon("", openFiles, func() {
		filereader.Filereader(a)
	})
	openFile := getIcon(IconStruct{source: "./images/file_icon.png"})
	lsCmd := widget.NewButtonWithIcon("LS", openFile, func() {
		w2 := a.NewWindow("file man")
		openFileFunc(w2, a)
	})
	office := getIcon(IconStruct{source: "./images/office.png"})
	officeBtn := widget.NewButtonWithIcon("office", office, func() {
		editeur.Editeur(a , "")
	})
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 100)), openBraveBtn, openCodeBtn, openFilesBtn, lsCmd, officeBtn)
	gridNewPos := canvas.NewPositionAnimation(fyne.NewPos(300, 600), fyne.NewPos(300, 600), time.Second, grid.Move)
	gridNewPos.Start()
	return grid
}
