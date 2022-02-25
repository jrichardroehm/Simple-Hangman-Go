package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func printAns(count int, ans [20]byte) {
	for i := 0; i < count; i++ {
		print(string(ans[i]))
	}
	print("\n")
}
func main() {
	var HangNoose [7]string
	HangNoose[0] = "--------\n |     |\n       |\n       |\n________"
	HangNoose[1] = "--------\n |     |\n O     |\n       |\n________"
	HangNoose[2] = "--------\n |     |\n O     |\n |     |\n________"
	HangNoose[3] = "--------\n |     |\n\\O     |\n |     |\n________"
	HangNoose[4] = "--------\n |     |\n\\O/    |\n |     |\n________"
	HangNoose[5] = "--------\n |     |\n\\O/    |\n |     |\n/_______"
	HangNoose[6] = "--------\n |     |\n\\O/    |\n |     |\n/ \\_____"

	var dict, err = os.ReadFile("Hangman\\hangDict.txt")
	//println(HangNoose[6])
	dict2 := strings.Split(string(dict[:]), "\n")
	//fmt.Println(len(dict))
	if err != nil {
		println(err)
	}
	//fmt.Println(rand.Intn(len(dict2) - 1))
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(dict2) - 1)
	answer := dict2[index]
	//ansBytes := [20]byte(answer)
	ansLen := len(answer)
	var guessArray [20]byte
	//println(ansLen)
	for i := 0; i < ansLen-1; i++ {
		guessArray[i] = '_'
		//print(string(guessArray[i]))
	}
	NooseCounter := 0
	println(HangNoose[NooseCounter])
	printAns(ansLen, guessArray)
	for answer[0:ansLen-1] != string(guessArray[:ansLen-1]) {

		var guess string
		fmt.Scan(&guess)
		if err != nil {
			print("Bad boy")
		}

		if strings.Contains(answer, guess) {
			for i := 0; i < ansLen; i++ {
				if (answer[i : i+1]) == guess {
					guessArray[i] = []byte(guess)[0]
					//println("In Loop #: ", NooseCounter)
				} else {
					//println("Jk its not in the loop?: ", guess, answer)
				}
			}
		} else {
			NooseCounter++
			if NooseCounter >= 7 {
				println("You looz lolz. Answer was: ", answer)
				return
			}
		}
		println(HangNoose[NooseCounter])
		printAns(ansLen, guessArray)
	}
	println("Congrats! You got the word right!")
	return
}

//Hang man
//Make dictionary of words;
//get random index for dictionary
//answer = dictionary(random)
//guessArray=[answer.length]
//for i = 0; i < answer.length;i++
//   guessArray[i]="_" <--- initialize the blanks for the hangman
//print HangNoose[0] version 1, 5 versions needed
//NooseCounter=0
//while answer != guessArray
//	userIn=getChar()
//	if(answer.contains(userIn))
//		index=answer.index(userIn)
//		guessArray[index]=userIn
//	else
//		NooseCounter++
//	if(NooseCounter=5)
//		print("You've failed. Answer was: ", answer,"\n")
//		quit
//	print(HangNoose[NooseCounter])
//	print(guessArray)
//print("Congrats! You're a genius!\n")
