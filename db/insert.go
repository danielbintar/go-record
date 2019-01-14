package db

import (
  "fmt"
  "github.com/danielbintar/go-record/structs"
  "github.com/danielbintar/go-record/connections"
)

func InsertQuery(model interface{}) string {
  table := structs.Init(model)

  query := "INSERT INTO `" + table.Name + "` ("
  counter := 0
  for _, column := range table.Columns {
    if column.Name == "id" {
      continue
    }
    if counter > 0 {
      query += ", "
    } else {
      counter++
    }
    query += "`" + column.Name + "`"
  }
  counter = 0
  query += ") VALUES ("
  for _, column := range table.Columns {
    if column.Name == "id" {
      continue
    }
    if counter > 0 {
      query += ", "
    } else {
      counter++
    }
    query += fmt.Sprintf("'%v'", column.Value)
  }
  query += ")"

  return query
}

func Insert(model interface{}) (interface{}, error) {
  q := InsertQuery(model)
  m,_ := connections.MySQLInstance()
  _, err := m.DB.Exec(q)
  return model, err
}
