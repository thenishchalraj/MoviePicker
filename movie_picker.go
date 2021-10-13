package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var i int
	fmt.Println("How many movies you want to enter:")
    	fmt.Scanf("%d", &i)

	movies := make([]string, i)
	for j := 0; j < i; j++ {
  		fmt.Println("Type ", j+1, " movie with '_' instead of space and hit enter")
		var movie string
		fmt.Scanf("%s", &movie)
		movies = append(movies, movie)
	}

	rand.Seed(time.Now().UnixNano())
	pickedMovie := movies[rand.Intn(len(movies))]

    	fmt.Println("Movie to watch:", pickedMovie)
}
