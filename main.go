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

	textcase := widget.NewEntry()
	textcase.SetText("textcase")

	genBtn := widget.NewButton("generate password", func() {
		pass := crypto.GenPass()
		textcase.SetText(pass)
		textcase.Disable()
	})

	var edit *widget.Button

	edit = widget.NewButton("edit", func() {
		if edit.Text == "edit" {
			textcase.Enable()
			edit.SetText("done")
		} else {
			textcase.Disable()
			edit.SetText("edit")
		}
	})

	w.SetContent(container.NewVBox(
		textcase,
		genBtn,
		edit,
	))

	w.Resize(fyne.NewSize(250.0, 120.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
