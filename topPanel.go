package main

import (
	"fmt"
	"pass-safe/crypto"
	"pass-safe/file"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func newTopPanel(w fyne.Window, mapData map[string][]string) *fyne.Container {
	search := widget.NewEntry()

	go func() {
		for ; ; time.Sleep(time.Second) {
			fmt.Println(search.Text)
			fmt.Println(mapData)
		}
	}()

	add := widget.NewButton("add", func() {
		username := widget.NewEntry()
		username.Validator = validation.NewRegexp(".+", "none")
		password := widget.NewPasswordEntry()
		password.Validator = validation.NewRegexp(".+", "none")
		suggest := widget.NewButton("suggest password", func() {
			password.SetText(crypto.GeneratePassword(16))
			password.Password = false
			password.Refresh()
		})

		items := []*widget.FormItem{
			widget.NewFormItem("Username", username),
			widget.NewFormItem("Password", password),
			widget.NewFormItem("", suggest),
		}

		dialog.ShowForm("Add a pair of new account and password", "OK", "Cancel", items, func(b bool) {
			if !b {
				return
			}

			if _, boolean := mapData[username.Text]; boolean {
				fmt.Println("duplicate")
				return
			}

			mapData[username.Text] = []string{password.Text}
			file.Store(mapData, key)
		}, w)
	})

	top := container.NewBorder(nil, nil, widget.NewLabel("search"), add, search)

	return top
}
