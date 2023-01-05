package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harish876/go_api/configs"
	"github.com/harish876/go_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	filter := bson.M{}
	findOptions := options.Find()
	cursor, err := userCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrases Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  users,
		"total": len(users),
	})

}
