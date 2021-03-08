package main

import (
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func showBody(w fyne.Window, data *storage.Data) {
	cont := container.NewGridWithColumns(3)
	grid := container.NewGridWithColumns(3)
	split := container.NewBorder(nil, nil, nil, grid, cont)
	scroll := container.NewVScroll(split)

	for i, v := range data.Map {
		name := widget.NewLabel(i)

		account := widget.NewEntry()
		account.Disable()
		account.SetText(v.Account)

		password := widget.NewEntry()
		password.Password = true
		password.SetText(v.Password)
		password.Disable()

		copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			w.Clipboard().SetContent(password.Text)
		})

		var edit *widget.Button

		edit = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
			if edit.Icon == theme.SettingsIcon() {
				password.Password = false
				account.Enable()
				password.Enable()
				edit.SetIcon(theme.ConfirmIcon())
			} else {
				password.Password = true
				account.Disable()
				password.Disable()
				edit.SetIcon(theme.SettingsIcon())

				data.Assign(name.Text, storage.NewPair(account.Text, password.Text))
			}
		})

		var del *widget.Button

		del = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			d := dialog.NewConfirm("Confirm", "do you want to delete?", func(b bool) {
				if b {
					data.Delete(name.Text)
					name.Hide()
					account.Hide()
					password.Hide()
					copyBtn.Hide()
					edit.Hide()
					del.Hide()
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

	top := newTopPanel(w, data, cont, grid)

	mainCont := container.NewBorder(top, nil, nil, nil, scroll)

	w.SetContent(mainCont)
}
