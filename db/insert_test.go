package db_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/danielbintar/go-record/db"
)

type willTransformInsertExample1 struct {
  Id        int
  Name      string
}

type willTransformInsertExample2 struct {
  Name      string
  Email     string
}

func TestInsertQueryString(t *testing.T) {
  example1 := &willTransformInsertExample1{
    Id: 0,
    Name: "lala",
  }

  example2 := &willTransformInsertExample2{
    Name: "lala",
    Email: "a@a.com",
  }

  assert.Equal(t, "INSERT INTO `will_transform_insert_example_1s` (`name`) VALUES ('lala')", db.InsertQuery(&example1))
  assert.Equal(t, "INSERT INTO `will_transform_insert_example_2s` (`name`, `email`) VALUES ('lala', 'a@a.com')", db.InsertQuery(&example2))

}
