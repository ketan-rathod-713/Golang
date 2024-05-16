package restapis

import (
	"context"
	"encoding/json"
	"graphql_search/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (a *api) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	authToken := r.URL.Query().Get("authToken")

	if authToken == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "Auth Token Not Provided."})
		return
	}

	userClaims, err := a.AuthService.VerifyJwt(authToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "Auth Token Not Provided."})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Check if email is already verified
	var userDB models.UserDB
	result := a.Database.Collection(a.DB_Collections.USERS).FindOne(ctx, bson.M{"_id": userClaims.ObjectId})
	err = result.Decode(&userDB)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error decoding userdb." + err.Error()})
		return
	}

	if userDB.IsEmailVerified {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "email is already verified"})
		return
	}

	updateValue := bson.M{
		"$set": bson.M{
			"isemailverified": true,
		},
	}
	// Not already verified hence mark it verified.
	_, err = a.Database.Collection(a.DB_Collections.USERS).UpdateOne(ctx, bson.M{"_id": userClaims.ObjectId}, updateValue)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error verifying email" + err.Error()})
		return
	}

	// mark email verified to true.
	userDB.IsEmailVerified = true
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&userDB)
}
