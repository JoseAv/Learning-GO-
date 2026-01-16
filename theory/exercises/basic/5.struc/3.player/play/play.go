package play

import "fmt"

type Player struct {
	id     int
	Life   float64
	points int
}

var CurrentId = 1

func AddPlayer(newPlayer Player, Players *[]Player) {
	newPlayer.id = CurrentId
	CurrentId++
	*Players = append(*Players, newPlayer)
	fmt.Printf("%+v\n", *Players)
}

func SumPoint(id int, addPoint int, Players *[]Player) {

	for i, val := range *Players {
		if val.id == id {
			(*Players)[i].points += addPoint
		}
	}

	fmt.Printf("%+v\n", *Players)
}



func ResPoint(id int, resPoint int, Players *[]Player) {
	for i, val := range *Players {
		if val.id == id {
			(*Players)[i].points -= resPoint
		}
	}
	
	fmt.Printf("%+v\n", *Players)
}


func MaxPoints(Players *[]Player) {
	if len(*Players) <=0  { 
		fmt.Println("No hay jugadores") 
		return
	 }
	 
	resultPlayer := (*Players)[0]
	for _, val := range *Players {
		if(resultPlayer.points < val.points){
			resultPlayer = val
		}
	}
	
	fmt.Printf("%+v\n", resultPlayer)
}