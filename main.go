package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
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
	// timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	var score int
	problems, err := problemPuller("quiz.csv")
	if err != nil {
		fmt.Println(err)
	}
	// loop over the problems and check answer to verify the score
	for i, p := range problems {
		fmt.Printf("%d. What is %s? ", i+1, p.Question)
		var ans string
		_, err := fmt.Scanf("%v", &ans)
		if err != nil {
			fmt.Println(err)
		}
		if ans == p.Answer {
			score++
		}

	}

	fmt.Println("Score was ", score)

}
