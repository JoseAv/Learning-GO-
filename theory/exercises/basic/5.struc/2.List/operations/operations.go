package operations

import "fmt"

type Works struct {
	title     string
	completed bool
}

func AddWork(title string, ListWorks *[]Works) []Works {
	newWork := Works{title: title, completed: true}
	*ListWorks = append(*ListWorks, newWork)
	return *ListWorks
}

func UpdateWork(index int, ListWorks *[]Works){
	if (index > len(*ListWorks)-1){
		return
	}
	work := &(*ListWorks)[index]
	work.completed = !work.completed
}


func FilterWork(state bool, ListWork *[]Works) []Works {
	newWork := []Works{}
	for _, val := range *ListWork {
		if state == val.completed {
			newWork = append(newWork, val)
		}
	}
	return newWork
}

func Resume(ListWork *[]Works) {
	for i, val := range *ListWork {
		fmt.Print(i)
		fmt.Printf("%+v\n", val)

	}
}