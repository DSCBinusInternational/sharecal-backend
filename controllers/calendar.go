package controllers

import (
	"context"
	"fmt"
	"net/http"
	"sharecal-backend/db"
	"sharecal-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCalendar(name string) (*models.Calendar, error) {
	var result models.Calendar
	mongoClient := db.GetMongo()
	calCollection := mongoClient.Database("sharecal").Collection("Calendar")
	if err := calCollection.FindOne(context.TODO(), bson.M{
		"_id": name,
	}).Decode(&result); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &result, nil
}

func AddCalendar(calName string, start string, end string, eventName string, notes string, pass string, color string) bool {
	mongoClient := db.GetMongo()
	calCollection := mongoClient.Database("sharecal").Collection("Calendar")

	// Parse the time to go object
	startDate, error := time.Parse(time.RFC3339, start)
	endDate, error2 := time.Parse(time.RFC3339, end)

	if error != nil {
		fmt.Println(error)
		return false
	}
	if error2 != nil {
		fmt.Println(error2)
		return false
	}

	// Create the netry object
	entry := models.TimeEntry{
		Name:  eventName,
		Time:  []time.Time{startDate, endDate},
		Notes: notes,
		Color: color,
	}

	// Try to find the document in mongo. If it does not exist, create it.
	var result bson.M
	if err := calCollection.FindOne(context.TODO(), bson.M{
		"_id": calName,
	}).Decode(&result); err != nil {
		// If the database is empty, insert it
		fmt.Printf("Calendar \"%s\" empty, creating", calName)
		calCollection.InsertOne(context.TODO(), models.Calendar{
			Id:       calName,
			Passcode: pass,
			Data: models.YearEntry{
				startDate.Year(): models.MonthEntry{
					int(startDate.Month()): models.DayEntry{
						startDate.Day(): []models.TimeEntry{
							entry,
						},
					},
				},
			},
		})
		return true
	}

	if result["passcode"] != pass {
		return false
	}

	// Gets the complete key to update the document
	dateKey := fmt.Sprintf("data.%d.%d.%d", startDate.Year(), int(startDate.Month()), startDate.Day())

	// Updates the existing document with a new entry
	calCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": calName},
		bson.D{
			{"$push", bson.M{dateKey: entry}},
		})
	return true
	// data, err3 := bson.Marshal(result)
	// if err3 != nil {
	// 	fmt.Println(err3.Error())
	// 	return
	// }
	// fmt.Println(data)
}

func CheckPasscode(name string, pass string) (bool, error) {
	var result bson.M
	mongoClient := db.GetMongo()
	calCollection := mongoClient.Database("sharecal").Collection("Calendar")
	// return pass == calCollection
	if err := calCollection.FindOne(context.TODO(), bson.M{
		"_id": name,
	}).Decode(&result); err != nil {
		fmt.Println(err)
		return false, err
	}
	return pass == result["passcode"], nil
}

func GetFunc(ctx *gin.Context) {
	name := ctx.Param("name")
	cal, err := GetCalendar(name)
	if err == nil {
		ctx.JSON(200, cal)
	} else {
		fmt.Println(err)
		ctx.JSON(404, gin.H{"success": false})
	}
}

func PostFunc(ctx *gin.Context) {
	body := models.CalendarEntry{}
	name := ctx.Param("name")

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if AddCalendar(name, body.Start, body.End, body.EventName, body.Notes, body.Pass, body.Color) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	} else {
		ctx.JSON(400, gin.H{
			"success": false,
		})
	}
}

func PassCheckFunc(ctx *gin.Context) {
	res, err := CheckPasscode(ctx.Param("name"), ctx.Param("pass"))
	if err == nil {
		ctx.JSON(200, gin.H{"success": true, "result": res})
	} else {
		ctx.JSON(404, gin.H{"success": false})
	}
}
