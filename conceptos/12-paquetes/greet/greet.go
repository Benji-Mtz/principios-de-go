package greet

var emoji = "😝"

// GreetEnglish retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
