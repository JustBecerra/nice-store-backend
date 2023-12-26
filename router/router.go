package router

import (
	"net/http"
	"strconv"

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

var products = []Product{
    {ID: 1, Title: "Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops", Price: 109.95, Description: "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday", Category: "men's clothing", Image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg", Rating: Rating{Rate: 3.9, Count: 120}},
    {ID: 2, Title: "Mens Casual Premium Slim Fit T-Shirts", Price: 22.3, Description: "Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.", Category: "men's clothing", Image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg", Rating: Rating{Rate: 4.1, Count: 259}},
    {ID: 3, Title: "Mens Cotton Jacket", Price: 55.99, Description: "great outerwear jackets for Spring/Autumn/Winter, suitable for many occasions, such as working, hiking, camping, mountain/rock climbing, cycling, traveling or other outdoors. Good gift choice for you or your family member. A warm-hearted love to Father, husband or son in this Thanksgiving or Christmas Day.", Category: "men's clothing", Image: "https://fakestoreapi.com/img/71li-ujtlUL._AC_UX679_.jpg", Rating: Rating{Rate: 4.7, Count: 500}},
}

// getAlbums responds with the list of all albums as JSON.
func getProducts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, products)
}

func getProductById(c *gin.Context) {
    id := c.Param("id")
	// convert id of type int to string
    idInt, err := strconv.Atoi(id)

    // Loop over the list of products, looking for
    // an products whose ID value matches the parameter.
    if(err != nil) {
        for _, a := range products {
            if a.ID == idInt {
                c.IndentedJSON(http.StatusOK, a)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
    }
    
}

func postProducts(c *gin.Context) {
    var newProduct Product

    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&newProduct); err != nil {
        return
    }

    // Add the new product to the slice.
    products = append(products, newProduct)
    c.IndentedJSON(http.StatusCreated, newProduct)
}

func InitRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/products", getProducts)
    // router.POST("/products", postProducts)
    router.GET("/products/:id", getProductById)
    return router
 }