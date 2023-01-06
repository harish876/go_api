package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harish876/go_api/configs"
	"github.com/harish876/go_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUser(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	filter := bson.M{"_id": objId}
	findResult := userCollection.FindOne(ctx, filter)
	if err = findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found here",
			"error":   err,
		})
	}

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found here",
			"error":   err,
		})
	}
	var user models.User
	findResult.Decode(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})

}

func CreateUser(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot create user-1",
			"error":   err,
		})
	}
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot create user-2",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User inserted successfully",
	})

}

func UpdateUser(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found here",
			"error":   err,
		})
	}
	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User payload incorrect",
			"error":   err,
		})
	}
	update := bson.M{"$set": updatedUser}
	filter := bson.M{"_id": objId}
	result, err := userCollection.UpdateOne(ctx, filter, update)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User updated successfully",
	})

}
func DeleteUser(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found here",
			"error":   err,
		})
	}
	filter := bson.M{"_id": objId}
	deletedUser, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found here",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    deletedUser,
		"success": true,
		"message": "User deleted successfully",
	})

}
