package main

import (
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var key = []byte{123, 4, 56, 98, 20, 14, 36, 44, 10, 255, 39, 223, 109, 182}

func main() {
	a := app.New()
	w := a.NewWindow("Password Safe")

	cont := container.NewGridWithColumns(3)
	grid := container.NewGridWithColumns(3)
	split := container.NewBorder(nil, nil, nil, grid, cont)
	scroll := container.NewVScroll(split)

	data := storage.NewData(key)
	data.Load()

	for i, v := range data.Map {
		name := widget.NewLabel(i)

		account := widget.NewEntry()
		account.Disable()
		account.SetText(v.Account)

		password := widget.NewEntry()
		password.Password = true
		password.SetText(v.Password)
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
				data.Delete(account.Text)
			} else {
				password.Password = true
				account.Disable()
				password.Disable()
				edit.SetText("edit")

				data.Assign(name.Text, storage.NewPair(account.Text, password.Text))
			}
		})

		del := widget.NewButton("delete", func() {
			d := dialog.NewConfirm("Confirm", "do you want to delete?", func(b bool) {
				if b {
					data.Delete(name.Text)
				}
			}, w)
			d.Show()
		})

		cont.Add(name)
		cont.Add(account)
		cont.Add(password)
		grid.Add(copyBtn)
		grid.Add(edit)
		grid.Add(del)
	}

	top := newTopPanel(w, data)

	mainCont := container.NewBorder(top, nil, nil, nil, scroll)

	w.SetContent(mainCont)

	w.Resize(fyne.NewSize(840.0, 600.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
