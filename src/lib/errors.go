package lib

import (
  "fmt"
)

func HandleError(err error) {
  if err != nil {
    fmt.Printf("Error: %s", err.Error)
  }
}