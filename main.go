package main

import (
	"pass-safe/crypto"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Password Safe")

	cont := container.NewGridWithColumns(3)

	for i := 0; i < 5; i++ {
		textcase := widget.NewEntry()
		textcase.Password = true
		textcase.SetText(crypto.GeneratePassword(16))

		copyBtn := widget.NewButton("copy", func() {
			w.Clipboard().SetContent(textcase.Text)
		})

		var edit *widget.Button

		edit = widget.NewButton("edit", func() {
			if edit.Text == "edit" {
				textcase.Password = false
				textcase.Enable()
				edit.SetText("done")
			} else {
				textcase.Password = true
				textcase.Disable()
				edit.SetText("edit")
			}
		})

		cont.Add(textcase)
		cont.Add(copyBtn)
		cont.Add(edit)
	}

	w.SetContent(cont)

	w.Resize(fyne.NewSize(450.0, 120.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
