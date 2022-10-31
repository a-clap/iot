package display

type ID uint
type Language uint

const (
	// PL is the only one supported (for now)
	PL Language = iota
)

const (
	Dark ID = iota
	Light
	MainWindow
	SettingsWindow
	size
)

var text = [...][size]string{
	PL: {
		Dark:           "Ciemny",
		Light:          "Jasny",
		MainWindow:     "Główne",
		SettingsWindow: "Ustawienia",
	},
}

var currentLanguage Language = PL

func SetLanguage(l Language) {
	currentLanguage = l
}

func Text(pos ID) string {
	return text[currentLanguage][pos]
}
