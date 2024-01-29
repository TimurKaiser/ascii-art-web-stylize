package controllers

import (
	"bufio"
	"os"
)

func getFont(font string) ([]string, error) {
	// Par exemple, chargez les lignes depuis un fichier ou une source de données.

	fileName := "static/font/" + font + ".txt" 
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
