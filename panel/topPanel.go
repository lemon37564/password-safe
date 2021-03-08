package panel

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

type topPanel struct {
	searchIcon *widget.Icon
	search     *widget.Entry
	add        *widget.Button
	data       *storage.Data
	middlePart *fyne.Container
	rightPart  *fyne.Container
}

func newTopPanel(w fyne.Window, data *storage.Data, middlePart *fyne.Container, rightPart *fyne.Container) *fyne.Container {

	top := topPanel{data: data, middlePart: middlePart, rightPart: rightPart}

	top.add = top.newAddBtn(w)
	top.search = top.newSearchEntry()
	top.searchIcon = top.newSearchIcon()

	topContainer := container.NewBorder(nil, nil, top.searchIcon, top.add, top.search)

	go func() {
		for ; ; time.Sleep(time.Second * 1) {
			fmt.Println(top.search.Text)
			fmt.Println(data)
		}
	}()

	return topContainer
}

func (t topPanel) newSearchIcon() *widget.Icon {
	return widget.NewIcon(theme.SearchIcon())
}

func (t topPanel) newSearchEntry() *widget.Entry {
	return widget.NewEntry()
}

func (t topPanel) newAddBtn(w fyne.Window) *widget.Button {
	add := widget.NewButtonWithIcon("add", theme.ContentAddIcon(), func() {
		notEmpty := validation.NewRegexp(".+", "cannot be none")

		name := widget.NewEntry()
		name.Validator = notEmpty

		account := widget.NewEntry()
		account.Validator = notEmpty

		password := widget.NewPasswordEntry()
		password.Validator = notEmpty

		suggest := widget.NewButton("suggest password", func() {
			password.Password = false
			password.SetText(crypto.GeneratePassword(16))
		})

		items := []*widget.FormItem{
			widget.NewFormItem("Name", name),
			widget.NewFormItem("Account", account),
			widget.NewFormItem("Password", password),
			widget.NewFormItem("", suggest),
		}

		dial := t.newDialog(w, items, name, account, password)
		dial.Show()
	})

	return add
}

func (t topPanel) newDialog(w fyne.Window, items []*widget.FormItem, name *widget.Entry, account *widget.Entry, password *widget.Entry) dialog.Dialog {
	d := dialog.NewForm("Add a pair of new account and password", "Confirm", "Cancel", items, func(boolean bool) {
		if !boolean {
			return
		}

		if _, boolean := t.data.Map[name.Text]; boolean {
			dialog.ShowError(errors.New("this name has existed"), w)
			return
		}
		t.data.Assign(name.Text, storage.NewPair(account.Text, password.Text))

		b := newBar(w, t.data, name.Text, account.Text, password.Text)

		t.middlePart.Add(b.nameCard)
		t.middlePart.Add(b.accountEntry)
		t.middlePart.Add(b.passwordEntry)
		t.rightPart.Add(b.copyButton)
		t.rightPart.Add(b.editButton)
		t.rightPart.Add(b.deleteButton)
	}, w)

	return d
}
