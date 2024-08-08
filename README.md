### Hangman Game with Golang

According to Wikipedia, Hangman is a guessing game for two or more players. One player thinks of a word, phrase, or sentence and the other(s) tries to guess it by suggesting letters or numbers within a certain number of guesses. Originally a paper-and-pencil game but I implemented it using Golang.

### Preconditions

* Golang should be installed at your PC. I used the version 1.22.6.

### Running the Code

Run ```go run .``` command in the root directory and the game will start. Then you should enter a single character to guess the word which is a movie name. If you can guess the word, you will win the game. Otherwise you will lose it after 9 attempts without success.