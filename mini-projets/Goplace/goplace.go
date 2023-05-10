package goplace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// param src: nom de fichier source
// param old: ancien mot
// param new: nouveau mot
// return occ: nombre d'occurences de old
// return lines: slice des numéros de lignes ou old a été trouvé
// return err: error de la fonction

// func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
// 	scanner := bufio.NewScanner("test.txt")
// 	fmt.Println(scanner)
// }

const (
	FILE_NAME string = "text.txt"
)

func FindReplaceFile(src, old, new string) (occ int, err error) {
	occurence := 0
	file, err := os.Open(src)

	defer file.Close()

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		textLowerCased := strings.ToLower(text)
		occurence += strings.Count(textLowerCased, old)
	}

	fmt.Println("occurence:", occurence)

	return occurence, nil
}

// param line: linge à traiter
// param old: ancien mot
// param new: nouveau mot
// return found: vrai si au moins une occurence est trouvée
// return result: résultat de remplacement (res == line)
// return occ: nombre d'occurence de old dans la ligne
// func ProcessLine(line, old, new string) (found bool, result string, occ int) {

// }
