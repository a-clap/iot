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

	split := container.NewHSplit(navigation(), view)
	split.Offset = 0.2

	w.SetContent(split)
	w.Resize(fyne.Size{
		Width:  display.Width(),
		Height: display.Height(),
	})

	w.ShowAndRun()
}

func navigation() fyne.CanvasObject {
	guiApp := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return display.ChildUIDs(uid)
		},
		IsBranch: func(uid string) bool {
			children := display.ChildUIDs(uid)

			return children != nil && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			// TODO: What is this?
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			//t, ok := gui.Windows[uid]
			//if !ok {
			//	fyne.LogError("Missing tutorial panel: "+uid, nil)
			//	return
			//}
			//obj.(*widget.Label).SetText(t.Name())
		},
		OnSelected: func(uid string) {
		},
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton(display.Text(display.Dark), func() {
			guiApp.Settings().SetTheme(display.DarkTheme())
		}),
		widget.NewButton(display.Text(display.Light), func() {
			guiApp.Settings().SetTheme(display.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
