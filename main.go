package main

import (
	"fmt"
	"pass-safe/crypto"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println(crypto.GeneratePassword(2048))
	return

	a := app.New()
	w := a.NewWindow("Password Safe")

	textcase := widget.NewEntry()
	textcase.SetText("textcase")

	genBtn := widget.NewButton("generate password", func() {
		pass := crypto.GeneratePassword(16)
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
	edit.Resize(edit.MinSize())

	cont := container.NewAdaptiveGrid(
		3,
		textcase,
		genBtn,
		edit,
	)

	w.SetContent(cont)

	w.Resize(fyne.NewSize(250.0, 120.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
