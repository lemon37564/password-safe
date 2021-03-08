package panel

import (
	"pass-safe/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type bar struct {
	nameCard      *widget.Card
	accountEntry  *widget.Entry
	passwordEntry *widget.Entry
	copyButton    *widget.Button
	editButton    *widget.Button
	deleteButton  *widget.Button
}

func newBar(w fyne.Window, data *storage.Data, name, account, password string) bar {
	b := bar{}

	b.nameCard = widget.NewCard("", name, nil)

	b.accountEntry = widget.NewEntry()
	b.accountEntry.Disable()
	b.accountEntry.SetText(account)

	b.passwordEntry = widget.NewPasswordEntry()
	b.passwordEntry.SetText(password)
	b.passwordEntry.Disable()

	b.copyButton = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(b.passwordEntry.Text)
	})

	b.editButton = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		b.passwordEntry.Password = !b.passwordEntry.Password
		if b.editButton.Icon == theme.SettingsIcon() {
			b.accountEntry.Enable()
			b.passwordEntry.Enable()
			b.editButton.SetIcon(theme.ConfirmIcon())
		} else {
			b.accountEntry.Disable()
			b.passwordEntry.Disable()
			b.editButton.SetIcon(theme.SettingsIcon())

			data.Assign(name, storage.NewPair(b.accountEntry.Text, b.passwordEntry.Text))
		}
	})

	b.deleteButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		d := dialog.NewConfirm("Confirm", "do you want to delete?", func(boolean bool) {
			if boolean {
				data.Delete(name)
				b.nameCard.Hide()
				b.accountEntry.Hide()
				b.passwordEntry.Hide()
				b.copyButton.Hide()
				b.editButton.Hide()
				b.deleteButton.Hide()
			}
		}, w)
		d.Show()
	})

	return b
}

func showBody(w fyne.Window, data *storage.Data) {
	middlePart := container.NewGridWithColumns(3)
	rightPart := container.NewGridWithColumns(3)
	split := container.NewBorder(nil, nil, nil, rightPart, middlePart)
	scroll := container.NewVScroll(split)

	for i, v := range data.GetMap() {
		b := newBar(w, data, i, v.Account, v.Password)

		middlePart.Add(b.nameCard)
		middlePart.Add(b.accountEntry)
		middlePart.Add(b.passwordEntry)
		rightPart.Add(b.copyButton)
		rightPart.Add(b.editButton)
		rightPart.Add(b.deleteButton)
	}

	top := newTopPanel(w, data, middlePart, rightPart)

	mainCont := container.NewBorder(top, nil, nil, nil, scroll)

	w.SetContent(mainCont)
}
