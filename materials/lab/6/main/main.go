package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"fmt"
	"log"
	"net/http"
	"scrape/scrape"

	"bufio"
	"os"

	"github.com/gorilla/mux"
)

//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

func main() {
	fmt.Println("When LOG_LEVEL = 0 DO NOT LOG anything")
	fmt.Println("When LOG_LEVEL = 1 LOG API details only")
	fmt.Println("When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)")

	//prompt for user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please select a log number based on the above options: ")
	text, _ := reader.ReadString('\n')

	fmt.Print(text)
	if text == "0\n" {
		scrape.Log_lvl = 0
		fmt.Println("Logging set to level 0\n\tWhen LOG_LEVEL = 0 DO NOT LOG anything")
	} else if text == "1\n" {
		scrape.Log_lvl = 1
		fmt.Println("Logging set to level 1\n\tWhen LOG_LEVEL = 1 LOG API details only")
	} else if text == "2\n" {
		scrape.Log_lvl = 2
		fmt.Println("Logging set to level 2\n\t2 LOG API details and file matches (e.g., everything)")
	} else {
		fmt.Println("NO valid level inputed, defaulting to 0")
		scrape.Log_lvl = 0
	}
	if scrape.Log_lvl > 0 {
		log.Println("starting API server")
	}

	//create a new router
	router := mux.NewRouter()
	if scrape.Log_lvl > 0 {
		log.Println("creating routes")
	}

	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	router.HandleFunc("/addsearch/{regex}", scrape.AddRegEx).Methods("GET")
	router.HandleFunc("/clear", scrape.ClearRegEx).Methods("GET")
	router.HandleFunc("/reset", scrape.ResetRegEx).Methods("GET")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
