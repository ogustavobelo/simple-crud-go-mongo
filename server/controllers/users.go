package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/models"
	"github.com/ogustavobelo/simple-crud-go/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// CreateUser godoc
// @Summary Create an user
// @ID post-create-user
// @Accept  json
// @Produce  json
// @Param data body models.CreateUserPayload true "Create User Payload"
// @Success 200 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		core.Error(c, "Can't create user!", err)
		return
	}

	if user.Email == "" {
		core.Error(c, "Email can't be empty!", nil)
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		core.Error(c, "Unable to save user!", err)
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID.String())
	if err != nil {
		core.UnknownError(c, err)
		return
	}

	core.Success(c, gin.H{
		"message": "user create sucessfully!",
		"user":    user,
		"token":   token,
	})
}

func GetUserByID(id string, c *gin.Context, ctx context.Context) {

	fmt.Println("Query ID ", id)
	parsedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		core.Error(c, "Invalid ID!", err)
		return
	}
	var user models.User
	filter := bson.D{{Key: "_id", Value: parsedID}}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		core.Error(c, "No user fonded!", err)
		return
	}

	core.Success(c, gin.H{
		"user": user,
	})
}

func ListUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	id := c.Query("id")
	if id != "" {
		GetUserByID(id, c, ctx)
	} else {
		ListAllUsers(c, ctx)
	}
}

func ListAllUsers(c *gin.Context, ctx context.Context) {

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		core.Error(c, "Can't list users!", err)
		return
	}
	defer cur.Close(ctx)
	users := []models.User{}
	for cur.Next(ctx) {
		var result models.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("user: ", result)
		users = append(users, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	core.Success(c, gin.H{
		"users": users,
	})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		core.Error(c, "Invalid user!", err)
		return
	}
	if user.ID == primitive.NilObjectID {
		core.Error(c, "User ID Can't be Empty!", nil)
		return
	}

	now := time.Now()
	user.UpdatedAt = &now

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := collection.FindOneAndUpdate(ctx,
		bson.M{"_id": user.ID},
		bson.M{
			"$set": user,
		})
	err := result.Err()
	if err != nil {
		core.Error(c, "Can't update user!", err)
		return
	}

	core.Success(c, gin.H{
		"message": "user updated succesfully",
		"user":    user,
	})
}

func DeleteAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection.DeleteMany(ctx, bson.M{})
}

func SetCollection(c *mongo.Collection) {
	collection = c
}
