package structs

import "github.com/danielbintar/go-record/normalizer"

type Column struct {
  StructName  string
  Name        string
  Value       interface{}
}

func CreateColumn(name string, value interface{}) *Column {
  return &Column{
    StructName: name,
      Name: normalizer.SafeName(name),
      Value: value,
  }
}

func FindColumnByName(columns []*Column, name string) *Column {
  for _, column := range columns {
    if column.Name == name {
      return column
    }
  }
  panic("column " + name + "not found")
}
