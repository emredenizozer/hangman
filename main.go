package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func main() {
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWords(targetWord)
	gameState := 0

	for !isGameOver(targetWord, guessedLetters, gameState) {
		printGameState(targetWord, guessedLetters, gameState)
		input := readInput()

		if len(input) != 1 {
			fmt.Println("Invalid input. Only a single character is allowed!")
			continue
		}

		letter := unicode.ToLower(rune(input[0]))
		if isCorrectGuess(targetWord, letter) {
			if guessedLetters[unicode.ToLower(letter)] {
				fmt.Printf("\n%c is already guessed. Please try an another character! \n\n", unicode.ToUpper(letter))
			} else {
				guessedLetters[unicode.ToLower(letter)] = true
			}
		} else {
			gameState++
		}
	}

	printGameState(targetWord, guessedLetters, gameState)

	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Printf("\nCongratulations! You guessed the word: %s\n", targetWord)
	} else if gameState >= 9 {
		fmt.Printf("\nSorry, you lost. The word was: %s\n", targetWord)
	} else {
		panic("\nInvalid game state. Exiting now.")
	}
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, gameState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || gameState >= 9
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}

	return true
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	gameState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(gameState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func getHangmanDrawing(gameState int) string {
	data, err := os.ReadFile(
		fmt.Sprintf("states/state%d", gameState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
