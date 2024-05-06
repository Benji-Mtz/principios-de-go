package main

import "fmt"

func main() {
	music := make(map[string]string)

	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	// Maps literales
	tech := map[string]string{
		"computador": "💻",
		"mouse":      "🖱️",
	}

	fmt.Println(tech)

	delete(tech, "computador")
	fmt.Println(tech)

	content, ok := music["fake"]
	fmt.Println(content, ok)

	content1, ok1 := music["guitar"]
	fmt.Println(content1, ok1)
}
