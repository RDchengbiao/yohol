package main

import (
	"./router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name string
	Sn   string
}

func TestFindBook(c *gin.Context) {
	sn, _ := c.GetQuery("sn")
	if sn == "" {
		sn = "1"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", sn, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("find book : %s", val))
}


func TestInsertBook(c *gin.Context) {
	sn, _ := c.GetQuery("sn")
	name, _ := c.GetQuery("name")
	res := "no insert"
	if sn != "" && name != "" {
		session, err := mgo.Dial("mongo_server")
		if err != nil {
			panic(err)
		} else {
			println("yyyyyyyyyyyyyyyyyyy")
		}
		collections := session.DB("test").C("book")

		book := &Book{
			Name: name,
			Sn:   sn,
		}
		err = collections.Insert(book)
		if err != nil {
			res = "no insert"
		}
		res = "insert!"
	}
	c.String(http.StatusOK, res)
}


func main() {
	r := gin.New()
	router.RegisterRoutes(r)
	r.Run(":8000")
	select {}
}