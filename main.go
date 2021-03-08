package main

import (
	"pass-safe/panel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Password Safe")

	panel.Startup(w)

	w.Resize(fyne.NewSize(800.0, 600.0))
	w.CenterOnScreen()
	w.ShowAndRun()
}
