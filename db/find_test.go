package db_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/danielbintar/go-record/db"
)

type willTransformExampleFindTest1 struct {
  Id        int
  Name      string
}

func TestFindQueryString(t *testing.T) {
  example1 := &willTransformExampleFindTest1{
    Id: 1,
    Name: "lala",
  }

  assert.Equal(t, "SELECT * FROM `will_transform_example_find_test_1s` WHERE `id` = 1 LIMIT 1;", db.FindQuery(&example1))
}

func TestFindByQueryString(t *testing.T) {
  example1 := &willTransformExampleFindTest1{
    Id: 1,
    Name: "lala",
  }

  options1 := []string{"username", "=", "lala"}
  options2 := []string{"login_count", "=", "5"}
  var options [][]string
  options = append(options, options1)
  assert.Equal(t, "SELECT * FROM `will_transform_example_find_test_1s` WHERE `username` = 'lala' LIMIT 1;", db.FindByQuery(&example1, options))

  options = append(options, options2)
  assert.Equal(t, "SELECT * FROM `will_transform_example_find_test_1s` WHERE `username` = 'lala' AND `login_count` = '5' LIMIT 1;", db.FindByQuery(&example1, options))
}
