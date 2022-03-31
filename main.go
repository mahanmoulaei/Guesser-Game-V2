package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func main() {
	http.HandleFunc("/ReceiveInput", receiveInput)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

/*
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
*/

type Result struct {
	GuessedNumberOne       int
	GuessedNumberTwo       int
	GuessedNumberThree     int
	GuessedNumberFour      int
	GuessedNumberFive      int
	GuessedNumbersString   string `json:"guessed_numbers"`
	GeneratedNumbersString string `json:"generated_numbers"`
	MatchNumbersCounter    int    `json:"match_counter"`
}

func receiveInput(w http.ResponseWriter, r *http.Request) {

	receivedJSON := r.URL.Query().Get("a")

	var guessedNumbersObject Result
	json.Unmarshal([]byte(receivedJSON), &guessedNumbersObject)

	generatedNumbers := GenerateRandomArrayOfNumbers(5, 0, 12)
	guessedNumbersObject.GeneratedNumbersString = GetArrayAsString(generatedNumbers)

	//Very Bad Way To Do This Shit
	var guessedNumbers []int
	guessedNumbers = append(guessedNumbers, guessedNumbersObject.GuessedNumberOne)
	guessedNumbers = append(guessedNumbers, guessedNumbersObject.GuessedNumberTwo)
	guessedNumbers = append(guessedNumbers, guessedNumbersObject.GuessedNumberThree)
	guessedNumbers = append(guessedNumbers, guessedNumbersObject.GuessedNumberFour)
	guessedNumbers = append(guessedNumbers, guessedNumbersObject.GuessedNumberFive)
	//End Of The Shitty Way
	guessedNumbersObject.GuessedNumbersString = GetArrayAsString(guessedNumbers)

	//GetMatchNumbers/Result
	guessedNumbersObject.MatchNumbersCounter = GetResult(guessedNumbers, generatedNumbers)

	out, err := json.MarshalIndent(guessedNumbersObject, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func GenerateRandomArrayOfNumbers(arrayLength int, min int, max int) []int {
	var array []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrayLength; i++ {
		for {
			var randomNumber = rand.Intn((max+1)-min) + min
			if !IsInArray(randomNumber, array) {
				array = append(array, randomNumber)
				break
			}
		}
	}
	return array
}

func IsInArray(element int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}
	return false
}

func GetArrayAsString(array []int) string {
	var output string
	var lengthOfArray int = len(array)
	for i := 0; i < lengthOfArray; i++ {
		if i != (lengthOfArray - 1) {
			output += strconv.Itoa(array[i]) + ", "
		} else {
			output += strconv.Itoa(array[i])
		}
	}
	return output
}

func GetResult(guessedNumbersArray []int, generatedNumbersArray []int) int {
	var counter int = 0
	var tempArray []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if generatedNumbersArray[j] == guessedNumbersArray[i] {
				if !IsInArray(guessedNumbersArray[i], tempArray) {
					tempArray = append(tempArray, guessedNumbersArray[i])
					counter++
				}
			}
		}
	}
	return counter
}
