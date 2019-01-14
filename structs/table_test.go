package structs_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/danielbintar/go-record/structs"
)

type willTransformExample1 struct {
  Id int
}

type willTransformExample2 struct {
  Id     int
  Name   string
}

type willTransformExample3 struct {
  Name   string
}

func TestInit(t *testing.T) {
  example1 := willTransformExample1{
    Id: 5,
  }

  example2 := willTransformExample2{
    Id: 7,
    Name: "lala",
  }

  example3 := willTransformExample3{
    Name: "lala",
  }

  assert.Equal(t, "will_transform_example_1s", structs.Init(example1).Name)
  assert.Equal(t, 1, len(structs.Init(example1).Columns))
  assert.Equal(t, "id", structs.Init(example1).Columns[0].Name)
  assert.Equal(t, 5, structs.Init(example1).Columns[0].Value)
  init_struct := structs.Init(example1)
  assert.Equal(t, 5, init_struct.Id())

  assert.Equal(t, "will_transform_example_2s", structs.Init(example2).Name)
  assert.Equal(t, 2, len(structs.Init(example2).Columns))
  assert.Equal(t, "id", structs.Init(example2).Columns[0].Name)
  assert.Equal(t, 7, structs.Init(example2).Columns[0].Value)
  assert.Equal(t, "name", structs.Init(example2).Columns[1].Name)
  assert.Equal(t, "lala", structs.Init(example2).Columns[1].Value)

  assert.Equal(t, "will_transform_example_3s", structs.Init(example3).Name)
  assert.Equal(t, 1, len(structs.Init(example3).Columns))
  assert.Equal(t, "name", structs.Init(example3).Columns[0].Name)
  assert.Equal(t, "lala", structs.Init(example3).Columns[0].Value)
}
