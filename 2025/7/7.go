package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {
		manifold := [][]rune{}

		shared.Open("7.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			manifold = append(manifold, []rune(line))
		})

		totalSplits := solve(manifold)
		return totalSplits
	})
}

func solve(manifold [][]rune) int {
	// Grille originale (pour l'affichage)
	original := manifold

	// Grille de travail pour les faisceaux
	beams := make([][]bool, len(manifold))
	for i := range manifold {
		beams[i] = make([]bool, len(manifold[i]))
	}

	// Trouver la position de départ 'S'
	startRow := -1
	for i := 0; i < len(original); i++ {
		for j := 0; j < len(original[i]); j++ {
			if original[i][j] == 'S' {
				startRow = i
				beams[i][j] = true // Le faisceau commence ici
				break
			}
		}
		if startRow != -1 {
			break
		}
	}

	totalSplits := 0

	// Affichage initial avant traitement
	displayManifold(original, beams, startRow-1, totalSplits)
	time.Sleep(300 * time.Millisecond)

	// Traitement récursif ligne par ligne
	for n := startRow; n < len(beams)-1; n++ {
		// Appliquer la transformation de la ligne n vers n+1
		splits := processLine(original, beams, n, n+1)
		totalSplits += splits

		// Afficher la visualisation (lignes n-5 à n+1)
		displayManifold(original, beams, n, totalSplits)
		time.Sleep(100 * time.Millisecond)
	}

	// Affichage final
	displayManifold(original, beams, len(beams)-1, totalSplits)
	fmt.Println()

	return totalSplits
}

// processLine traite la ligne n et met à jour la ligne n+1
// Retourne le nombre de splits effectués
func processLine(original [][]rune, beams [][]bool, n, nPlus1 int) int {
	if nPlus1 >= len(beams) {
		return 0
	}

	splits := 0
	width := len(beams[n])

	// Créer un tableau temporaire pour les nouveaux faisceaux
	nextBeams := make(map[int]bool)

	// Phase 1: Traiter tous les faisceaux descendants
	for col := 0; col < width; col++ {
		if beams[n][col] {
			// Un faisceau descend
			if original[nPlus1][col] == '^' {
				// Le faisceau rencontre un splitter
				splits++

				// Créer deux nouveaux faisceaux à gauche et à droite
				if col > 0 {
					nextBeams[col-1] = true
				}
				if col < width-1 {
					nextBeams[col+1] = true
				}
			} else {
				// Le faisceau continue (case vide ou déjà un faisceau)
				nextBeams[col] = true
			}
		}
	}

	// Phase 2: Appliquer les changements à la ligne n+1
	for col := 0; col < width; col++ {
		if nextBeams[col] {
			beams[nPlus1][col] = true
		}
	}

	return splits
}

// displayManifold affiche les lignes n-5 à n+1
func displayManifold(original [][]rune, beams [][]bool, currentN int, totalSplits int) {
	fmt.Print("\033[H\033[2J") // Nettoyer l'écran

	startLine := currentN - 5
	if startLine < 0 {
		startLine = 0
	}
	endLine := currentN + 6
	if endLine >= len(original) {
		endLine = len(original) - 1
	}
	if currentN < 0 {
		endLine = 5
		if endLine >= len(original) {
			endLine = len(original) - 1
		}
	}

	output := strings.Builder{}
	output.WriteString("╔════════════════════════════════════╗\n")
	if currentN < 0 {
		output.WriteString("║  Tachyon Manifold - INITIAL STATE ║\n")
	} else {
		output.WriteString(fmt.Sprintf("║  Tachyon Manifold - Line %3d/%3d  ║\n", currentN, len(original)-1))
	}
	output.WriteString(fmt.Sprintf("║         Total Splits: %3d          ║\n", totalSplits))
	output.WriteString("╠════════════════════════════════════╣\n")

	for i := startLine; i <= endLine; i++ {
		output.WriteString("║ ")
		for j := 0; j < len(original[i]); j++ {
			c := original[i][j]
			hasBeam := beams[i][j]

			// Afficher selon le type de cellule et la présence de faisceau
			if c == 'S' {
				if hasBeam {
					output.WriteString("\033[1;32mS\033[0m") // Vert gras pour S avec faisceau
				} else {
					output.WriteString("\033[32mS\033[0m") // Vert pour S
				}
			} else if c == '^' {
				if hasBeam {
					output.WriteString("\033[1;31m|\033[0m") // Rouge gras pour faisceau sur splitter
				} else {
					output.WriteString("\033[31m^\033[0m") // Rouge pour splitter non touché
				}
			} else if hasBeam {
				output.WriteString("\033[33m|\033[0m") // Jaune pour faisceau dans le vide
			} else {
				output.WriteRune(c) // Caractère original
			}
		}
		if i == currentN {
			output.WriteString(" \033[36m← n\033[0m")
		} else if i == currentN+1 {
			output.WriteString(" \033[35m← n+1\033[0m")
		}
		output.WriteString("\n")
	}

	output.WriteString("╚════════════════════════════════════╝\n")
	fmt.Print(output.String())
}
