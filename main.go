package main

import (
	"encoding/json"
	"fmt"
	"sharecal-backend/models"
)

// getCalendar - return a
// func getCalendar(calName string) {

// }

// func addCalendar(start time.Date, end time.Date, eventName string, notes string, calName string, pass string) {
// 	http.HandleFunc("/encode", func(w http.ResponseWriter, r *htpp.Request) {
// 		event := Event{
// 			name:    eventName,
// 			time[0]: start,
// 			time[1]: end,
// 			notes:   notes
// 		}
// 		json.NewEncoder(w).Encode(event)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// func checkPasscode(calName Calendar.Bruhcal, pass string) {
// 	if pass == calName.Passcode {
// 		fmt.Println("Pass is correct")
// 	} else {
// 		fmt.Println("Pass is incorrect")
// 	}
// }

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

	var data models.Calendar
	var err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data.Data["2022"]["5"]["7"])
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
	fmt.Println(data.Passcode)
	// checkPasscode("bruhcal", "rwfnwxe")
}
