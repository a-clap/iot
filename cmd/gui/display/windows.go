package display

import (
	"errors"
	"fyne.io/fyne/v2"
)

var (
	ErrNotExist = errors.New("window doesn't exist")
)

func WelcomeWindow() Window {
	return mainMenu
}

func GetWindow(uid string) (Window, error) {
	w, ok := windows[uid]
	if !ok {
		return nil, ErrNotExist
	}
	return w, nil
}

type Window interface {
	key() string
	Title() string
	ChildUIDs() []string
	Selected() fyne.CanvasObject
	UnSelected()
}

// windowNav sets first navigation view
type windowNav struct{}

var _ Window = windowNav{}

func (n windowNav) ChildUIDs() []string {
	return []string{mainMenu.key(), settingsMenu.key()}
}

var (
	navMenu      = windowNav{}
	mainMenu     = &windowMain{}
	settingsMenu = &windowSettings{}
)

var windows = map[string]Window{
	navMenu.key():      navMenu,
	mainMenu.key():     mainMenu,
	settingsMenu.key(): settingsMenu,
}

func (n windowNav) key() string {
	return ""
}

func (n windowNav) Title() string {
	return ""
}

func (n windowNav) Selected() fyne.CanvasObject {
	return nil
}

func (n windowNav) UnSelected() {
}
