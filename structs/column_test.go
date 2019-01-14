package structs_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/danielbintar/go-record/structs"
)

type columnExample struct {
  name  string
  value interface{}
}

func TestCreateColumn(t *testing.T) {
  example1 := columnExample{
    name: "Id",
    value: 7,
  }

  assert.Equal(t, "id", structs.CreateColumn(example1.name, example1.value).Name)
  assert.Equal(t, "Id", structs.CreateColumn(example1.name, example1.value).StructName)
  assert.Equal(t, 7, structs.CreateColumn(example1.name, example1.value).Value)
}

func TestFindColumnByName(t *testing.T) {
  example1 := columnExample{
    name: "Id",
    value: 7,
  }

  example2 := columnExample{
    name: "Lala",
    value: 7,
  }

  firstColumn := structs.CreateColumn(example1.name, example1.value)
  secondColumn := structs.CreateColumn(example2.name, example2.value)

  arr := []*structs.Column{firstColumn, secondColumn}

  assert.Equal(t, firstColumn, structs.FindColumnByName(arr, "id"))
  assert.Equal(t, secondColumn, structs.FindColumnByName(arr, "lala"))
}
