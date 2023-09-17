package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quantity"`
}

var books =[]book{
	{ID:"1",Title: "Anna Karenina",Author: "Leo Tolstoy",Quantity: 10},
	{ID:"2",Title: "Pride and Prejudice",Author: "Jane Austen",Quantity: 68},
	{ID:"3",Title: "Arthashastra",Author: "Kautilya",Quantity: 19},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id,ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing Id Query Parameter."})
		return
	}
	book,err := getBookById(id)

	if err!= nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not available."})
		return
	}

	book.Quantity -=1

	c.IndentedJSON(http.StatusOK,book)

}

func checkInBook(c *gin.Context) {
	id,ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing Id Query Parameter."})
		return
	}
	book,err := getBookById(id)

	if err!= nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found."})
		return
	}


	book.Quantity +=1

	c.IndentedJSON(http.StatusOK,book)

}


func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err!= nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found."})
		return
	}
	
	c.IndentedJSON(http.StatusOK,book)
}

func getBookById(id string) (*book,error) {
	for i, b:=range books {
		if b.ID == id {
			return &books[i],nil
		}
	}
	return nil,errors.New("books not found")
}

func AddBooks(c *gin.Context){
	var newBook book
	if err:= c.BindJSON(&newBook); err !=nil{
		return
	}
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func main(){
	router := gin.Default()	
	router.GET("/books",getBooks)
	router.GET("/book/:id",bookById)
	router.POST("/AddBook",AddBooks)
	router.PATCH("/checkoutBook",checkoutBook)
	router.PATCH("/checkinBook",checkInBook)
	router.Run("localhost:8080")

}