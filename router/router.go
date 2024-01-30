package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"store/nice-store-backend/db"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Rating struct {
    Rate float64 `json:"rate"`
    Count int `json:"count"`
}

type Product struct {
    ID     int  `json:"id"`
    Title  string  `json:"title"`
    Price  float64 `json:"price"`
	Description string `json:"description"`
	Category string `json:"category"`
    Image string `json:"image"`
    Rating Rating `json:"rating"`
}

type User struct {
    ID int `json:"id"`
    Fullname string `json:"fullname"`
    Email string `json:"email"`
    Password string `json:"password"`
    Address string `json:"address"`
    // Image image.Image `json:"image"`
}

// getAlbums responds with the list of all albums as JSON.
func getProducts(c *gin.Context) {
    response, err := http.Get("https://fakestoreapi.com/products")

    if err != nil {
        fmt.Print(err.Error())
    }
    defer response.Body.Close()

    // read response
    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Print(err)
    }

    //unmarshal to transform into JSON
    var responseObject []Product
    json.Unmarshal(responseData, &responseObject)
    err = json.Unmarshal(responseData, &responseObject)
    if err != nil {
        fmt.Print(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response"})
        return
    }

    // send unmarshal response
    c.IndentedJSON(http.StatusOK, responseObject)
}

func getProductById(c *gin.Context) {
    id := c.Param("id")
	// convert id of type int to string
    idInt, err := strconv.Atoi(id)
    url := fmt.Sprintf("https://fakestoreapi.com/products/%d", idInt)
    response, err := http.Get(url)

    if err != nil {
        fmt.Print(err.Error())
    }
    defer response.Body.Close()
    
    // read response
    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Print(err)
    }

    //unmarshal to transform into JSON
    var responseObject Product
    json.Unmarshal(responseData, &responseObject)
    err = json.Unmarshal(responseData, &responseObject)
    if err != nil {
        fmt.Print(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response"})
        return
    }

    // send unmarshal response
    c.IndentedJSON(http.StatusOK, responseObject)
    
}

func updateUser(c *gin.Context) {
    db := db.GetDB()

    var userToUpdate User
    var bodyData User
    if err := c.ShouldBindJSON(&bodyData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

    matchID := bodyData.ID
    fmt.Println(matchID)

    err := db.First(&userToUpdate, matchID).Error
    if err != nil {
        // Handle the error, you can send an error response or log it
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
        return
    }

    if bodyData.Fullname != "" {
		userToUpdate.Fullname = bodyData.Fullname
	}
	if bodyData.Email != "" {
		userToUpdate.Email = bodyData.Email
	}
	if bodyData.Password != "" {
		userToUpdate.Password = bodyData.Password
	}
	if bodyData.Address != "" {
		userToUpdate.Address = bodyData.Address
	}

    err = db.Save(&userToUpdate).Error
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User could not be updated"})
		return 
	}

    // Send a success response if the update was successful
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func InitRouter() *gin.Engine {
    router := gin.Default()
    db.GetDB()
    router.Use(cors.Default())
    router.GET("/products", getProducts)
    router.GET("/products/:id", getProductById)
    router.PUT("/user", updateUser)
    return router
 }