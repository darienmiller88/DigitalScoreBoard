package services

import(
	"net/http"
	"fmt"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

//This handler will allow the client to change the name of a particular player
func UpdatePlayerName(req *http.Request, locationName string, oldPlayerName string, newPlayerName string) models.Result[*mongo.UpdateResult]{
	updateUserResult := models.Result[*mongo.UpdateResult]{
		StatusCode: http.StatusInternalServerError,
	}
	
	//Filter first for the location, and then for the player in the "users" array field.
	filter := bson.M{ "location_name": locationName, "users.name": oldPlayerName }
	update := bson.M{ "$set": bson.M{ "users.$.name": newPlayerName } }

	//Using the above filters, find the specific user in the array at the specific location, and change their name.
	updateOneResult, err := database.GetLocationsCollection().UpdateOne(req.Context(), filter, update)

	if err != nil{
		updateUserResult.Err = err

		return updateUserResult
	}

	//If the update was succesful, return an Ok(200) and the update return data.
	updateUserResult.ResultData = updateOneResult
	updateUserResult.StatusCode = http.StatusOK

	return updateUserResult
}

//Service function to allow adding or removing a user to or from an Adapt location.
func UpdatePlayersForLocation(req *http.Request, mongoUpdateOperator string, locationName string, username string) models.Result[*mongo.UpdateResult] {
	updateUserResult := models.Result[*mongo.UpdateResult]{
		StatusCode: http.StatusInternalServerError,
	}
	
	if !(mongoUpdateOperator == PULL || mongoUpdateOperator == PUSH){
		updateUserResult.Err = fmt.Errorf("mongo operator %s not valid. Must be either %s or %s", mongoUpdateOperator, PULL, PUSH)

		return updateUserResult
	}

	filter := bson.M{"location_name": locationName}
	update := bson.M{mongoUpdateOperator: bson.M{"users": bson.M{"name": username}}}
 	updateOneResult, err := database.GetLocationsCollection().UpdateOne(req.Context(), filter, update)

	if err != nil{
		updateUserResult.Err = err
		
		return updateUserResult
	}

	if updateOneResult.ModifiedCount == 0 {
		updateUserResult.Err = fmt.Errorf("no location \"%s\" found OR player \"%s\" not found", locationName, username)
		updateUserResult.StatusCode = http.StatusNotFound

		return updateUserResult
	}

	updateUserResult.ResultData = updateOneResult
	updateUserResult.StatusCode = http.StatusOK

	return updateUserResult
}