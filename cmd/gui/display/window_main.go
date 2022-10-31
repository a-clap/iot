package display

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

type windowMain struct {
}

var _ Window = &windowMain{}

func (w *windowMain) key() string {
	return "windowMain"
}

func (w *windowMain) Title() string {
	return Text(MainWindow)
}
func (w *windowMain) ChildUIDs() []string {
	return nil
}

func (w *windowMain) Selected() fyne.CanvasObject {
	return widget.NewButton("hello button", func() {
		fmt.Println("tapped")
	})
}

func (w *windowMain) UnSelected() {
	log.Println("UnSelected")
}
