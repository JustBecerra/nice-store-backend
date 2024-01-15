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
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
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

func CreateUser(user *User) (*User, error) {
    database := db.GetDB()
    // user.ID = uuid.New().String()

    res := database.Create(user)
    if res.Error != nil {
        return nil, res.Error
    }
    return user, nil
 }

 func postUser(ctx *gin.Context) {
    var user User
    err := ctx.Bind(&user)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    res, err := CreateUser(&user)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{
        "movie": res,
    })
 }

func InitRouter() *gin.Engine {
    router := gin.Default()
    db.GetDB()
    router.Use(cors.Default())
    router.GET("/products", getProducts)
    router.GET("/products/:id", getProductById)
    router.POST("/user", postUser)
    return router
 }