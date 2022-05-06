package main

import (
	"fmt"
	"sharecal-backend/controllers"
	"sharecal-backend/db"
)

func main() {

	db.Init()

	// controllers.AddCalendar("StephCal", "2022-05-03T10:00:00+07:00", "2022-05-07T14:30:00+07:00", "Halloween", "Rememeber to dress up", "limatigapuluh", "#000000")
	// controllers.AddCalendar("StephCal", "2022-05-03T15:00:00+07:00", "2022-05-07T18:30:00+07:00", "Go Home", "Pick up Anne", "limatigapuluh", "#ff00ff")
	fmt.Println(controllers.CheckPasscode("StephCal", "rwfnwxe"))
	fmt.Println(controllers.CheckPasscode("StephCal", "limatigapuluh"))
	fmt.Println(controllers.GetCalendar("StephCal"))
}
