package goplace

import (
	"bufio"
	"os"
	"strings"
)

const (
	FILE_NAME string = "text.txt"
)

// param src: nom de fichier source
// param old: ancien mot
// param new: nouveau mot
// return occ: nombre d'occurences de old
// return lines: slice des numéros de lignes ou old a été trouvé
// return err: error de la fonction
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	wordOccurence := 0
	fileLines := 0
	wordOccurenceLines := []int{}

	textLowerCased := ""
	file, err := os.Open(src)

	if err != nil {
		return 0, wordOccurenceLines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLines++
		text := scanner.Text()

		// managing new lines
		textLowerCased += text + "\n"
		wordOccurence += strings.Count(textLowerCased, old)

		// getting the number of occurence lines
		if strings.Contains(strings.ToLower(text), old) {
			wordOccurenceLines = append(wordOccurenceLines, fileLines)
		}
	}

	return wordOccurence, wordOccurenceLines, nil
}
