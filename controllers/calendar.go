package controllers

import (
	"context"
	"fmt"
	"sharecal-backend/db"
	"sharecal-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// getCalendar - return a
// func getCalendar(calName string) {
// }

func AddCalendar(calName string, start string, end string, eventName string, notes string, pass string, color string) {
	mongoClient := db.GetMongo()
	calCollection := mongoClient.Database("sharecal").Collection("Calendar")

	dateFormat := "2022-05-05T20:30:00+07:00"
	startDate, error := time.Parse(dateFormat, start)
	endDate, error2 := time.Parse(dateFormat, end)

	if error != nil {
		fmt.Println(error)
		return
	}
	if error2 != nil {
		fmt.Println(error)
		return
	}

	calObj := models.Calendar{
		Passcode: pass,
		Data: models.YearEntry{
			startDate.Year(): models.MonthEntry{
				int(startDate.Month()): models.DayEntry{
					startDate.Day(): []models.TimeEntry{
						{
							Name:  eventName,
							Time:  []time.Time{startDate, endDate},
							Notes: notes,
							Color: color,
						},
					},
				},
			},
		},
		Id: calName,
	}

	var result bson.M
	if err := calCollection.FindOne(context.TODO(), bson.M{
		"_id": calName,
	}).Decode(&result); err != nil {
		// If the database is empty, insert it
		fmt.Printf("Calendar %s empty, creating", calName)
		calCollection.InsertOne(context.TODO(), calObj)
		return
	}

	calCollection.UpdateOne(context.TODO(), bson.M{"_id": calName}, calObj)
}

func CheckPasscode(calendar models.Calendar, pass string) bool {
	return pass == calendar.Passcode
}
