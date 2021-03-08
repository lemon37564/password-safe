package panel

import (
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Startup(w fyne.Window) {

	var key string
	var data *storage.Data

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

			key = password.Text
			data = storage.NewData([]byte(key))

			err := data.Load()
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			w.SetOnClosed(data.Store)

			showBody(w, data)
		}, w)
	}
}

func show(w fyne.Window) {
	password := widget.NewPasswordEntry()
	password.Validator = validation.NewRegexp(".+", "none")
	confirm := widget.NewPasswordEntry()
	confirm.Validator = validation.NewRegexp(".+", "none")
	items := []*widget.FormItem{
		widget.NewFormItem("Password", password),
		widget.NewFormItem("Confirm", confirm),
	}

	var dial dialog.Dialog

	dial = dialog.NewForm("creating a new file, please enter your password:", "OK", "Cancel", items, func(b bool) {
		if !b {
			return
		}

		if password.Text != confirm.Text {
			err := dialog.NewInformation("not same", "not same", w)
			err.Show()
			return
		}

		storage.Create()
	}, w)

	dial.Show()
}
