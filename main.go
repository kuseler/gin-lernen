package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userWithoutPass struct {
	Username string `json:"username"`
}

// Method to return User without Password
func (u User) WithoutPassword() userWithoutPass {
	return userWithoutPass{
		Username: u.Username,
	}
}

type Recipe struct {
	ID      int    `json:"id"`
	Creator string `json:"creator"` // Refers to the `username` in users table
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func init() {
	// Connect to PostgreSQL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://user:password@localhost:5432/api/v2?sslmode=disable" // Replace with your credentials
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to verify database connection: %v", err)
	}
}

func main() {
	router := gin.Default()

	api := router.Group("/api/v2")
	{
		api.GET("/users/all", getAllUsers)
		api.POST("/users/register", registerUser)
		api.POST("/users/login", loginUser)
		api.GET("/recipes/all", getAllRecipes)
		api.POST("/recipes/create", createRecipe)
		api.GET("/users/allsecrets", getAllUserssecrets)
	}

	if err := router.Run(":9999"); err != nil {
		log.Fatalf("Failed to run server on port 9999: %v", err)
	}
	router.Static("/static", "./static") // Serves files in the "static" folder

}

// Handlers

func getAllUserssecrets(c *gin.Context) {
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username); err != nil {
			log.Printf("Error scanning user row: %v", err)
			continue
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func getAllUsers(c *gin.Context) {
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var usersWithoutPass []userWithoutPass
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username); err != nil {
			log.Printf("Error scanning user row: %v", err)
			continue
		}
		usersWithoutPass = append(usersWithoutPass, user.WithoutPassword())
	}

	c.JSON(http.StatusOK, usersWithoutPass)
}

func registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Insert user into the database
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := db.Exec(query, newUser.Username, newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginUser(c *gin.Context) {
	var loginDetails User
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", loginDetails.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows || storedPassword != loginDetails.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during authentication"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func getAllRecipes(c *gin.Context) {
	rows, err := db.Query("SELECT id, creator, title, content FROM recipes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var recipes []Recipe
	for rows.Next() {
		var recipe Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Creator, &recipe.Title, &recipe.Content); err != nil {
			log.Printf("Error scanning recipe row: %v", err)
			continue
		}
		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

func createRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if the creator (username) exists
	var creatorExists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)", newRecipe.Creator).Scan(&creatorExists)
	if err != nil || !creatorExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Creator username does not exist"})
		return
	}

	// Insert recipe into the database
	query := "INSERT INTO recipes (creator, title, content) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(query, newRecipe.Creator, newRecipe.Title, newRecipe.Content).Scan(&newRecipe.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Recipe created successfully", "recipe_id": newRecipe.ID})
}
