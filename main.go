package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string   `json:"username"`
	Recipes  []Recipe `json:"recipes"`
	Password string   `json:"password"`
}

// Users sind zusammen mit Rezepten, sollte separat,

type UserWithoutPass struct {
	Username string   `json:"username"`
	Recipes  []Recipe `json:"recipes"`
}

func (u User) WithoutPassword() UserWithoutPass {
	return UserWithoutPass{
		Username: u.Username,
		Recipes:  u.Recipes,
	}
}

type Recipe struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var users = []User{
	{Username: "john_doe", Recipes: []Recipe{{Title: "Pasta", Content: "Boil pasta, add sauce."}}},
}
var recipes = []Recipe{
	{Title: "Pasta", Content: "Boil pasta, add sauce."},
}

func main() {
	router := gin.Default()

	api := router.Group("/api/v2")
	{
		api.GET("/users/all", getAllUsers)
		api.POST("/users/register", registerUser)
		api.POST("/users/log-in", loginUser)
		api.GET("/users/:id", getUserByID)
		api.PUT("/users/:id", updateUserByID)
		api.DELETE("/users/:id", deleteUserByID)

		api.GET("/recipes/all", getAllRecipes)
		api.GET("/recipes/:id", getRecipeByID)
		api.POST("/recipes/:id", createRecipe)
		api.PUT("/recipes/:id", updateRecipeByID)
		api.DELETE("/recipes/:id", deleteRecipeByID)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

// Handlers
func getAllUsers(c *gin.Context) {
	var usersWithoutPassword []UserWithoutPass
	for _, user := range users {
		usersWithoutPassword = append(usersWithoutPassword, user.WithoutPassword())
	}
	c.JSON(http.StatusOK, usersWithoutPassword)
}

func registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginUser(c *gin.Context) {
	var loginDetails User
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, user := range users {
		if user.Username == loginDetails.Username {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.Username == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func updateUserByID(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.Username == id {
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}

func deleteUserByID(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.Username == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}

func getAllRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func getRecipeByID(c *gin.Context) {
	id := c.Param("id")
	for _, recipe := range recipes {
		if recipe.Title == id {
			c.JSON(http.StatusOK, recipe)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
}

func createRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipes = append(recipes, newRecipe)
	c.JSON(http.StatusCreated, gin.H{"message": "Recipe created successfully"})
}

func updateRecipeByID(c *gin.Context) {
	id := c.Param("id")
	var updatedRecipe Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, recipe := range recipes {
		if recipe.Title == id {
			recipes[i] = updatedRecipe
			c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfully"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}

func deleteRecipeByID(c *gin.Context) {
	id := c.Param("id")
	for i, recipe := range recipes {
		if recipe.Title == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
