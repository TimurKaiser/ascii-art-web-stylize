package controllers

import (
	"strings"
)

func convertInput(line string, lines []string) string {
	arr := strings.Split(line, "\n")

	var result strings.Builder
	
	if line != "" {
		for _, str := range arr {
			arr2 := strings.Split(str, "\\n")
			for _, str2 := range arr2 {
				for i := 1; i < 9; i++ {
					for _, char := range str2 {
						ascii := int(char) - 32
						index := ascii*9 + i

						if index >= 0 && index < len(lines) {
							result.WriteString(lines[index])
						} else {
							// Ignorer l'index invalide et ne rien imprimer dans le terminal
						}
					}
					result.WriteString("\n")
				}
				result.WriteString("\n")
			}
		}
	}

	return result.String()
}
