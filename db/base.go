package db

import (
  "strconv"
  "time"
  "reflect"

  "github.com/danielbintar/go-record/structs"
)

func RowToStruct(model interface{}, column_names []string, column_values []interface{}) {
  ps := reflect.ValueOf(model).Elem().Elem()
  table := structs.Init(model)

  for i, name := range column_names {
    structName := structs.FindColumnByName(table.Columns, name).StructName
    f := ps.FieldByName(structName)
    columnType := f.Type()

    switch columnType.String() {
    case "string":
      f.Set(reflect.ValueOf(string(column_values[i].([]byte))))
    case "int":
      v, err := strconv.Atoi(string(column_values[i].([]byte)))
      if err != nil { panic(err) }
      f.Set(reflect.ValueOf(v))
    case "time.Time":
      f.Set(reflect.ValueOf(column_values[i].(time.Time)))
    default:
      panic(columnType.String())
    }
  }
}
