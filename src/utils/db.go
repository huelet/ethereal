package utils

import (
  "os"
  
  "github.com/joho/godotenv"
  "github.com/huelet/ethereal/src/lib"
  "github.com/fauna/faunadb-go/v4/faunadb"
)

func InitDB() *faunadb.FaunaClient {
  err := godotenv.Load()
  lib.HandleError(err)
  dbKey := os.Getenv("FAUNA_DB_KEY")

  client := faunadb.NewFaunaClient(
    dbKey,
    faunadb.Endpoint("https://db.fauna.com/"),
  )
  
  return client
}