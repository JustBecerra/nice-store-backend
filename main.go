package main

import (
    "net/http"
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

func main() {
    router := gin.Default()
    router.GET("/products", getProducts)

    router.Run("localhost:8080")
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