package main

import (
	"encoding/json"
	"fmt"
	"sharecal-backend/controllers"
	"sharecal-backend/db"
	"sharecal-backend/models"
)

func main() {
	var jsonData = `{
		"passcode": "limatigapuluh",
		"data": {
			"2022": {
				"5": {
					"5": [
					{
						"name": "Make backend time",
						"time": ["2022-05-05T20:30:00+07:00", "2022-05-05T22:30:00+07:00"],
						"notes": "Made in golang",
						"color": "#ff0000"
					},
					{
						"name": "Sleep",
						"time": ["2022-05-05T23:45:00+07:00", "2022-05-06T08:30:00+07:00"],
						"notes": "Yes",
						"color": "#ff00ff"
					}
					],
					"7": [{
					"name": "React Last Day",
					"time": ["2022-05-07T13:00:00+07:00", "2022-05-07T14:30:00+07:00"],
					"notes": "Finally ;---;",
					"color": "#00ff00"
					}]
				}
			}
		}
	}`

	db.Init()

	var data models.Calendar
	var err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	controllers.AddCalendar("StephCal", "2022-05-07T13:00:00+07:00", "2022-05-07T14:30:00+07:00", "Halloween", "Rememeber to dress up", "sixtynine", "#000000")

	// fmt.Println(data.Data["2022"]["5"]["7"])
	// bs, _ := json.Marshal(data)
	// fmt.Println(string(bs))
	// fmt.Println(data.Passcode)
	// fmt.Println(controllers.CheckPasscode(data, "rwfnwxe"))
	// fmt.Println(controllers.CheckPasscode(data, "limatigapuluh"))
}
