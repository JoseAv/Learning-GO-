package main

import (
	"player/play"
)

func main() {

	Players:= []play.Player{}
	
	

	play.AddPlayer(play.Player{Life:200},&Players)
	play.AddPlayer(play.Player{Life:200},&Players)
	play.AddPlayer(play.Player{Life:200},&Players)

	play.SumPoint(1,100,&Players)
	play.SumPoint(2,200,&Players)
	play.SumPoint(3,150,&Players)
	
	play.ResPoint(1,100,&Players)
	play.ResPoint(2,200,&Players)
	play.ResPoint(3,150,&Players)

	play.MaxPoints(&Players)

}