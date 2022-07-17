package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func problemPuller(fileName string) ([]Problem, error) {
	b, err := ioutil.ReadFile("./" + fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading file")
	}
	var problems []Problem
	for _, line := range strings.Split(string(b), "\n") {
		val := strings.Split(line, ",")
		prob := Problem{Question: val[0], Answer: val[1]}
		problems = append(problems, prob)
	}
	return problems, nil
}
func main() {
	timer := flag.Int("t", 10, "timer for the quiz")

	flag.Parse()

	var score int
	problems, err := problemPuller("quiz.csv")
	if err != nil {
		fmt.Println(err)
	}
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	// loop over the problems and check answer to verify the score

problemLoop:
	for i, p := range problems {
		fmt.Printf("%d. What is %s? ", i+1, p.Question)
		var ans string
		go func() {
			fmt.Scanf("%v", &ans)
			ansC <- ans
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			fmt.Println("Time is up")
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.Answer {
				score++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}

	}

	fmt.Println("Score was ", score, "out of ", len(problems))

}
