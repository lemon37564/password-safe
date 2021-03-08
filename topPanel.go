package main

import (
	"errors"
	"fmt"
	"pass-safe/crypto"
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func newTopPanel(w fyne.Window, data *storage.Data) *fyne.Container {
	search := widget.NewEntry()

	go func() {
		// for ; ; time.Sleep(time.Second) {
		// 	fmt.Println(search.Text)
		// 	fmt.Println(data)
		// }
	}()

	add := widget.NewButton("add", func() {
		name := widget.NewEntry()
		name.Validator = validation.NewRegexp(".+", "none")
		account := widget.NewEntry()
		account.Validator = validation.NewRegexp(".+", "none")
		password := widget.NewPasswordEntry()
		password.Validator = validation.NewRegexp(".+", "none")
		suggest := widget.NewButton("suggest password", func() {
			password.SetText(crypto.GeneratePassword(16))
			password.Password = false
			password.Refresh()
		})

		items := []*widget.FormItem{
			widget.NewFormItem("name", name),
			widget.NewFormItem("account", account),
			widget.NewFormItem("Password", password),
			widget.NewFormItem("", suggest),
		}

		dialog.ShowForm("Add a pair of new account and password", "Confirm", "Cancel", items, func(b bool) {
			if !b {
				return
			}

			if _, boolean := data.Map[name.Text]; boolean {
				fmt.Println("duplicate")
				dialog.ShowError(errors.New("this name has existed"), w)
				return
			}

			data.Assign(name.Text, storage.NewPair(account.Text, password.Text))
		}, w)
	})

	top := container.NewBorder(nil, nil, widget.NewLabel("search"), add, search)

	return top
}
