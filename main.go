package main

import (
	"cloud/database/db"
	"github.com/gin-gonic/gin"
	"log"
)

const dbPath string = "./kvDB"

func main() {

	// database connection
	kvStore, err := db.NewKVStore(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer kvStore.Close()

	router := gin.Default()

	router.GET("/ping", HandlePing())
	router.POST("/set", HandleSet(kvStore))
	router.GET("/get", HandleGet(kvStore))

	router.Run("localhost:8000")
}
