package main

import (
	"fmt"

	"github.com/RoxyBodySlam/cinema/movie"
	"github.com/RoxyBodySlam/cinema/ticket"
)

/*
//package moview
func reviewMovie(name string, rating float64){
	fmt.Printf("I reviewed %s and it's rating is %f\n",name,rating)
}


func findMovieName(imdbID string) string{
	switch imdbID{
	case "tt4154796" :
		return "Avengers : Endgame"
	case "tt1825683" :
		return "Black Panther"
	}
	return "not found."
 }


 //package ticket
 func buyTicket(movie string){
	fmt.Printf("I bought tickets to %s\n",movie)
 }

*/

func inin() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)

}
