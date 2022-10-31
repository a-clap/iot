package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

func DarkTheme() fyne.Theme {
	return newTheme(theme.VariantDark)
}

func LightTheme() fyne.Theme {
	return newTheme(theme.VariantLight)
}

type themeVariant struct {
	v fyne.ThemeVariant
	fyne.Theme
}

func newTheme(v fyne.ThemeVariant) *themeVariant {
	return &themeVariant{
		v:     v,
		Theme: theme.DefaultTheme(),
	}
}

func (t *themeVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, t.v)
}
