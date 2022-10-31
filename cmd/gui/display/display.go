package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func Width() float32 {
	return 1024
}

func Height() float32 {
	return 768
}

func DarkTheme() fyne.Theme {
	return newTheme(theme.VariantDark)
}

func LightTheme() fyne.Theme {
	return newTheme(theme.VariantLight)
}

func ChildUIDs(uid string) []string {
	return []string{"1", "2", "3"}
}
