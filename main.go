package main

import (
	"pass-safe/file"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var key = []byte{123, 4, 56, 98, 20, 14, 36, 44, 10, 255, 39, 223, 109, 182}

func main() {
	a := app.New()
	w := a.NewWindow("Password Safe")

	cont := container.NewGridWithColumns(2)
	grid := container.NewGridWithColumns(2)
	split := container.NewBorder(nil, nil, nil, grid, cont)
	scroll := container.NewVScroll(split)

	mapData := file.Read(key)

	for i, v := range mapData {
		account := widget.NewEntry()
		account.Disable()
		account.SetText(i)

		password := widget.NewEntry()
		password.Password = true
		password.SetText(v[0])
		password.Disable()

		copyBtn := widget.NewButton("copy", func() {
			w.Clipboard().SetContent(password.Text)
		})

		var edit *widget.Button

		edit = widget.NewButton("edit", func() {
			if edit.Text == "edit" {
				password.Password = false
				account.Enable()
				password.Enable()
				edit.SetText("done")
				delete(mapData, account.Text)
			} else {
				password.Password = true
				account.Disable()
				password.Disable()
				edit.SetText("edit")

				mapData[account.Text] = []string{password.Text}

				file.Store(mapData, key)
			}
		})

		cont.Add(account)
		cont.Add(password)
		grid.Add(copyBtn)
		grid.Add(edit)
	}

	top := newTopPanel(w, mapData)

	mainCont := container.NewBorder(top, nil, nil, nil, scroll)

	w.SetContent(mainCont)

	w.Resize(fyne.NewSize(800.0, 600.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
