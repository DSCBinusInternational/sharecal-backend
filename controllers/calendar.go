package controllers

import (
	"context"
	"fmt"
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
	// return pass == calCollection
	if err := calCollection.FindOne(context.TODO(), bson.M{
		"_id": name,
	}).Decode(&result); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &result, nil
}

func AddCalendar(calName string, start string, end string, eventName string, notes string, pass string, color string) {
	mongoClient := db.GetMongo()
	calCollection := mongoClient.Database("sharecal").Collection("Calendar")

	// Parse the time to go object
	startDate, error := time.Parse(time.RFC3339, start)
	endDate, error2 := time.Parse(time.RFC3339, end)

	if error != nil {
		fmt.Println(error)
		return
	}
	if error2 != nil {
		fmt.Println(error2)
		return
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
		return
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
	ctx.JSON(200, gin.H{
		"message":"Hai JaON",
	})
	fmt.Println(name)
}

func PostFunc(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message":"Huha",
	})
}

func PutFunc(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message":"heoho",
	})
}

func HeloFunc(c *gin.Context){
	c.JSON(200, gin.H{"status":"helo"})
}
