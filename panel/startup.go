package panel

import (
	"errors"
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Startup construct whole window
func Startup(w fyne.Window) {

	if !storage.IsFileExist() {
		show(w)

	} else {
		password := widget.NewPasswordEntry()
		password.Validator = validation.NewRegexp(".+", "none")

		items := []*widget.FormItem{
			widget.NewFormItem("Password", password),
		}

		dialog.ShowForm("please enter password to continue:", "Confirm", "Cancel", items, func(b bool) {
			if !b {
				return
			}

			next(w, password.Text)
		}, w)
	}
}

func show(w fyne.Window) {
	notEmpty := validation.NewRegexp(".+", "none")

	password := widget.NewPasswordEntry()
	password.Validator = notEmpty
	confirm := widget.NewPasswordEntry()
	confirm.Validator = notEmpty

	items := []*widget.FormItem{
		widget.NewFormItem("Password", password),
		widget.NewFormItem("Confirm", confirm),
	}

	var login, err dialog.Dialog

	err = dialog.NewError(errors.New("two passwords are not the same"), w)
	err.SetOnClosed(func() { login.Show() })

	login = dialog.NewForm("creating a new file, please enter your password:", "OK", "Cancel", items, func(b bool) {
		if !b {
			return
		}

		if password.Text != confirm.Text {
			err.Show()
			return
		}

		storage.Create()
		next(w, password.Text)
	}, w)

	login.Show()
}

func next(w fyne.Window, key string) {
	data := storage.NewData([]byte(key))

	err := data.Load()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	w.SetOnClosed(data.Store)
	showBody(w, data)
}
