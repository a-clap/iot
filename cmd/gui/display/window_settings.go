package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type windowSettings struct {
}

var _ Window = &windowSettings{}

func (w *windowSettings) key() string {
	return "windowSettings"
}

func (w *windowSettings) Title() string {
	return Text(SettingsWindow)
}

func (w *windowSettings) ChildUIDs() []string {
	return nil
}

func (w *windowSettings) Selected() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("tab111111111111111111111111111", theme.MoveDownIcon(), widget.NewLabel("vey long word on safasfasfsa")),
		container.NewTabItemWithIcon("tab2", theme.MoveUpIcon(), widget.NewLabel("world")),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

func (w *windowSettings) UnSelected() {
}
