package connections

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
  "time"
  "os"
)

type MySQL struct {
  DB *sql.DB
}

type MySQLOption struct {
  User     string
  Password string
  Host     string
  Port     string
  Database string
  Charset  string
}

func MySQLInstance() (*MySQL, error) {
  opt := MySQLOption{
    User:     os.Getenv("MYSQL_USER"),
    Password: os.Getenv("MYSQL_PASSWORD"),
    Host:     os.Getenv("MYSQL_HOST"),
    Port:     os.Getenv("MYSQL_PORT"),
    Database: os.Getenv("MYSQL_DATABASE"),
    Charset:  os.Getenv("MYSQL_CHARSET"),
  }

  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset))
  if err != nil {
    return &MySQL{}, err
  }

  db.SetConnMaxLifetime(time.Second * 20)
  db.SetMaxIdleConns(0)

  return &MySQL{DB: db}, nil
}
