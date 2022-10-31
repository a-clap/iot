package display

type ID uint

const (
	Dark ID = iota
	Light
)

var text = [...]string{
	Dark:  "Ciemny",
	Light: "Jasny",
}

func Text(pos ID) string {
	return text[pos]
}
