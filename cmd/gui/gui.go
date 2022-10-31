package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/a-clap/iot/cmd/gui/display"
)

func main() {
	guiApp := app.NewWithID("dest.io")
	guiApp.Settings().SetTheme(display.DarkTheme())

	w := guiApp.NewWindow("app")

	content := container.NewMax()

	view := container.NewBorder(nil, nil, nil, nil, content)

	split := container.NewHSplit(navigation(content), view)
	split.Offset = 0.2

	w.SetContent(split)

	w.Resize(fyne.Size{
		Width:  display.Width(),
		Height: display.Height(),
	})

	w.ShowAndRun()
}

func navigation(c *fyne.Container) fyne.CanvasObject {

	tree := &widget.Tree{
		ChildUIDs: func(uid widget.TreeNodeID) []string {
			if w, err := display.GetWindow(uid); err == nil {
				return w.ChildUIDs()
			}
			return nil
		},
		IsBranch: func(uid widget.TreeNodeID) bool {
			if w, err := display.GetWindow(uid); err == nil {
				return len(w.ChildUIDs()) > 0
			}
			return false
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("")
		},
		UpdateNode: func(uid widget.TreeNodeID, branch bool, obj fyne.CanvasObject) {
			if w, err := display.GetWindow(uid); err == nil {
				obj.(*widget.Label).SetText(w.Title())
			}
		},
		OnSelected: func(uid widget.TreeNodeID) {
			if w, err := display.GetWindow(uid); err == nil {
				c.Objects = []fyne.CanvasObject{w.Selected()}
				c.Refresh()
			}
		},
		OnUnselected: func(uid widget.TreeNodeID) {
			if w, err := display.GetWindow(uid); err == nil {
				w.UnSelected()
			}
		},
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton(display.Text(display.Dark), func() {
			fyne.CurrentApp().Settings().SetTheme(display.DarkTheme())
		}),
		widget.NewButton(display.Text(display.Light), func() {
			fyne.CurrentApp().Settings().SetTheme(display.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
