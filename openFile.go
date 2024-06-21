package main

import (
	"fmt"
	"io/ioutil"

	"time"

	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/RazanakotoMandresy/os-fyne-go/internal/editeur"
)

var str = "/"
var foldPlace = ""
var theCont = ""
var boolsHide = true

// TO do trouver des bug et les corriger
func openFileFunc(w2 fyne.Window, a fyne.App) {
	var LsArr = []string{}
	grid := container.New(layout.NewGridLayout(4))
	g2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 100)))
	cmdLs := exec.Command("ls", str)
	var out strings.Builder
	cmdLs.Stdout = &out
	err := cmdLs.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	splitedLs := strings.Split(out.String(), "\n")
	LsArr = append(LsArr, splitedLs...)
	bac := backFunction(foldPlace, w2)
	g2.Add(bac)
	for _, lsArrRanged := range LsArr {
		fileIcon := iconTodisp(lsArrRanged)
		if lsArrRanged == "" || strings.Contains(lsArrRanged, "efi") || strings.Contains(lsArrRanged, "bin") || strings.Contains(lsArrRanged, "root") {
			continue
		}
		bools := false
		buttonName := nbCharactere(lsArrRanged)
		button := widget.NewButtonWithIcon(buttonName, fileIcon, func() {
			bools = !bools
			if bools {
				if strings.Contains(lsArrRanged, ".png") || strings.Contains(lsArrRanged, ".jpg") || strings.Contains(lsArrRanged, ".svg") || strings.Contains(lsArrRanged, ".jpeg") {
					displayImgages(a, str+"/"+lsArrRanged)
				} else if strings.Contains(lsArrRanged, "txt") || strings.Contains(lsArrRanged, "docx") || strings.Contains(lsArrRanged, "odt") || strings.Contains(lsArrRanged, ".doc") {
					content, err := ioutil.ReadFile(str + "/" + lsArrRanged)
					if err != nil {
						fmt.Println(err)
					}
					editeur.Editeur(a, string(content))
				} else {
					strp := &str
					*strp = str + "/" + lsArrRanged
					foldPlaceP := &foldPlace
					*foldPlaceP = lsArrRanged
					time.Sleep(time.Millisecond)
					openFileFunc(w2, a)
				}
			}
		})
		buttonForSpace := buttonSpace(lsArrRanged)
		cadre := widget.NewCard("", "", container.New(layout.NewVBoxLayout(), button, buttonForSpace))
		grid.Add(cadre)
	}
	left := allHardDrive(w2)
	info := infoSpace()
	info.Hidden = boolsHide
	cont := container.NewHBox(container.NewVBox(g2, left), grid, info)
	n := container.NewScroll(cont)
	w2.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeySpace {
			info := infoSpace()
			boolsHideP := &boolsHide
			*boolsHideP = !boolsHide
			info.Hidden = boolsHide
			time.Sleep(time.Millisecond)
			openFileFunc(w2, a)
		}
	})
	w2.SetContent(n)
	w2.Resize(fyne.NewSize(720, 720))
	w2.Show()
}
func iconTodisp(l string) *fyne.StaticResource {
	if strings.Contains(l, ".png") || strings.Contains(l, ".jpg") || strings.Contains(l, ".svg") || strings.Contains(l, ".jpeg") {
		folderIcon := getIcon(IconStruct{source: "./images/file_type_image.svg"})
		return folderIcon
	} else if strings.Contains(l, ".go") {
		folderIcon := getIcon(IconStruct{source: "./images/file_type_go.svg"})
		return folderIcon
	} else if strings.Contains(l, ".pptx") || strings.Contains(l, ".words") || strings.Contains(l, ".xlxs") {
		folderIcon := getIcon(IconStruct{source: "./images/office.png"})
		return folderIcon
	} else {
		folderIcon := getIcon(IconStruct{source: "./images/folder_icon.png"})
		return folderIcon
	}

}
func backFunction(lsArrRanged string, w2 fyne.Window) *widget.Button {
	backButton := widget.NewButton("retours", func() {
		orginale, _, _ := strings.Cut(str, lsArrRanged)
		strp := &str
		*strp = orginale
		fmt.Println(str)
		time.Sleep(time.Millisecond)
		openFileFunc(w2, a)

	})
	return backButton
}

func nbCharactere(lsArrRanged string) string {
	n := strings.Count(lsArrRanged, "")
	if n > 30 {
		lenStr := len(lsArrRanged) / 7
		part1 := lsArrRanged[:lenStr]
		doted := part1 + "..."
		return doted
	} else if n > 20 {
		lenStr := len(lsArrRanged) / 3
		part1 := lsArrRanged[:lenStr]
		doted := part1 + "..."
		return doted
	} else if n > 10 {
		lenStr := len(lsArrRanged) / 2
		part1 := lsArrRanged[:lenStr]
		doted := part1 + "..."
		return doted
	} else {
		return lsArrRanged
	}
}
