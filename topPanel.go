package main

import (
	"errors"
	"fmt"
	"pass-safe/crypto"
	"pass-safe/storage"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func newTopPanel(w fyne.Window, data *storage.Data, cont *fyne.Container, grid *fyne.Container) *fyne.Container {
	search := widget.NewEntry()

	go func() {
		for ; ; time.Sleep(time.Second) {
			fmt.Println(search.Text)
			fmt.Println(data)
		}
	}()

	add := widget.NewButtonWithIcon("add", theme.ContentAddIcon(), func() {
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
			widget.NewFormItem("Name", name),
			widget.NewFormItem("Account", account),
			widget.NewFormItem("Password", password),
			widget.NewFormItem("", suggest),
		}

		dialog.ShowForm("Add a pair of new account and password", "Confirm", "Cancel", items, func(b bool) {
			if !b {
				return
			}

			if _, boolean := data.Map[name.Text]; boolean {
				dialog.ShowError(errors.New("this name has existed"), w)
				return
			}
			data.Assign(name.Text, storage.NewPair(account.Text, password.Text))

			name := widget.NewLabel(name.Text)
			account1 := widget.NewEntry()
			account1.Disable()
			account1.SetText(account.Text)

			password1 := widget.NewEntry()
			password1.Password = true
			password1.SetText(password.Text)
			password1.Disable()

			copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
				w.Clipboard().SetContent(password1.Text)
			})

			var edit *widget.Button

			edit = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
				if edit.Icon == theme.SettingsIcon() {
					password1.Password = false
					account1.Enable()
					password1.Enable()
					edit.SetIcon(theme.ConfirmIcon())
				} else {
					password1.Password = true
					account1.Disable()
					password1.Disable()
					edit.SetIcon(theme.SettingsIcon())

					data.Assign(name.Text, storage.NewPair(account1.Text, password1.Text))
				}
			})

			var del *widget.Button

			del = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				d := dialog.NewConfirm("Confirm", "do you want to delete?", func(b bool) {
					if b {
						data.Delete(name.Text)
						name.Hide()
						account1.Hide()
						password1.Hide()
						copyBtn.Hide()
						edit.Hide()
						del.Hide()
					}
				}, w)
				d.Show()
			})

			cont.Add(name)
			cont.Add(account1)
			cont.Add(password1)
			grid.Add(copyBtn)
			grid.Add(edit)
			grid.Add(del)
		}, w)
	})

	top := container.NewBorder(nil, nil, widget.NewIcon(theme.SearchIcon()), add, search)

	return top
}
