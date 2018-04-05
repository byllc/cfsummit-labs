package main

import (
  "github.com/nanobox-io/golang-scribble"
  "github.com/google/uuid"
  "fmt"
  "time"
)

type ToDo struct {
  UID       string      `json:"id"`
  Name      string      `json:"name"`
  Created   string      `json:"created"`
  Done      string      `json:"done"`
}

func (toDo *ToDo) getToDo(db *scribble.Driver) {
  db.Read("toDo", toDo.UID, &toDo)
}
func (toDo *ToDo) putToDo(db *scribble.Driver) {
  if (toDo.UID == ""){
    toDo.UID = uuid.Must(uuid.NewRandom()).String()
  }
  db.Write("toDo", toDo.UID, toDo)
}
func (toDo *ToDo) finish(db *scribble.Driver) {
  if (toDo.UID == ""){
    fmt.Println("This todo cannot be finished as it does not exist yet")
  }
  toDo.Done=time.Now().String()
  db.Write("toDo", toDo.UID, toDo)
}
