package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	file, err := os.Open("./answer_key.csv")
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var correctAnswers int
	
	for _, pair := range records {
		var answer string
		clearTerminal()
		fmt.Println(pair[0])
		timer := time.NewTimer(3 * time.Second)
		defer timer.Stop()
		
		receivedAnswer := make(chan bool)
		
		go func() {
			_, err := fmt.Scan(&answer)
			if err != nil {
				log.Fatal(err)
			}
			receivedAnswer <- true
		}()
		
		select {
		case <-timer.C:
			clearTerminal()
			fmt.Println("Time up!")
			goto end
		case <-receivedAnswer:
			if(answer == pair[1]) {
				correctAnswers++
			}
		}
	}

	end:
	endgameMessage := fmt.Sprintf("You got %d correct answers out of %d", correctAnswers, len(records))
	fmt.Println(endgameMessage)
}