package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Password Safe")

	startup(w)

	w.Resize(fyne.NewSize(840.0, 600.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
