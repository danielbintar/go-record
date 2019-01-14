package structs

import (
  "reflect"
  "github.com/danielbintar/go-record/normalizer"
)

type Table struct {
  Name    string
  Columns []*Column
}

func (t *Table) Id() int {
  for _, column := range t.Columns {
    if column.Name == "id" {
      return column.Value.(int);
    }
  }
  return 0
}

func Init(model interface{}) Table {
  modelName := reflect.TypeOf(model).Elem().Elem().Name()
  tableName := normalizer.SafeName(modelName) + "s"

  v := reflect.ValueOf(model).Elem().Elem()
  columns := make([]*Column, v.NumField())
  for i := 0; i < v.NumField(); i++ {
    columns[i] = CreateColumn(v.Type().Field(i).Name, v.Field(i).Interface())
  }

  return Table{
    Name: tableName,
    Columns: columns,
  }
}
