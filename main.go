package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const DEFAULT_LENGTH int = 20
const MAX_LENGTH int = 99

func main() {
	a := app.New()
	w := a.NewWindow("Password Generator")

	title := canvas.NewText("Password length:", color.White)
	title.Alignment = fyne.TextAlignCenter

	input := widget.NewEntry()
	input.Text = strconv.Itoa(DEFAULT_LENGTH)
	input.Validator = lengthValidator

	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	btn := widget.NewButton("Generate password", func() {
		err := input.Validate()
		if err != nil {
			text.Text = err.Error()
		} else {
			passwordLength, _ := strconv.Atoi(input.Text)
			text.Text = GeneratePassword(passwordLength)
			w.Clipboard().SetContent(text.Text)
		}
		text.Refresh()
		resizeWindow(w)
	})

	content := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		title,
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), input),
		text,
		btn,
		layout.NewSpacer())

	w.SetContent(content)
	resizeWindow(w)
	w.ShowAndRun()
}

func resizeWindow(w fyne.Window) {
	w.Resize(fyne.NewSize(375, 200))
}

// returns nil if provided length is within boundaries
// or an error if this is not the case.
func lengthValidator(s string) error {
	len, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if len < 0 || len > MAX_LENGTH {
		return fmt.Errorf("password length must be between 1 and %d", MAX_LENGTH)
	}
	return nil
}
