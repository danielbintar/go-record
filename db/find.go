package db

import (
  "errors"
  "reflect"
  "strconv"
  "github.com/danielbintar/go-record/structs"
  "github.com/danielbintar/go-record/connections"
)

func FindQuery(model interface{}) string {
  table := structs.Init(model)
  query := "SELECT * FROM `" + table.Name + "` WHERE `id` = " + strconv.Itoa(table.Id()) + " LIMIT 1;"
  return query
}

func FindByQuery(model interface{}, options [][]string) string {
  table := structs.Init(model)
  query := "SELECT * FROM `" + table.Name + "` WHERE `"

  for i, option := range options {
    if i > 0 {
      query += " AND `"
    }
    query += option[0] + "` " + option[1] + " '" + option[2] + "'"
  }

  query += " LIMIT 1;"
  return query
}

func Find(model interface{}) error {
  q := FindQuery(model)

  columns_count := reflect.ValueOf(model).Elem().Elem().NumField()

  values := make([]interface{}, columns_count)
  columns := make([]interface{}, columns_count)

  for i, _ := range columns {
    columns[i] = &values[i]
  }

  m, err := connections.MySQLInstance()
  if err != nil { panic(err) }

  rows, err := m.DB.Query(q)
  if err != nil { panic(err) }

  var column_names []string

  defer rows.Close()
  if rows.Next() {
    column_names, err = rows.Columns()
    if err != nil { panic(err) }

    err = rows.Scan(columns...)
    if err != nil { panic(err) }
  } else {
    panic("no result for " + q)
  }

  RowToStruct(model, column_names, values)
  return err
}

func FindBy(model interface{}, options ...[]string) error {
  q := FindByQuery(model, options)

  columns_count := reflect.ValueOf(model).Elem().Elem().NumField()

  values := make([]interface{}, columns_count)
  columns := make([]interface{}, columns_count)

  for i, _ := range columns {
    columns[i] = &values[i]
  }

  m, err := connections.MySQLInstance()
  if err != nil { panic(err) }

  rows, err := m.DB.Query(q)
  if err != nil { panic(err) }

  var column_names []string

  defer rows.Close()
  if rows.Next() {
    column_names, err = rows.Columns()
    if err != nil { panic(err) }

    err = rows.Scan(columns...)
    if err != nil { panic(err) }
  } else {
    err = errors.New("not found")
  }

  if err != nil {
    return err
  }

  RowToStruct(model, column_names, values)
  return err
}
