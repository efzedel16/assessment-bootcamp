package config

import (
  "FauzulPasswordManager/migrate"
  "fmt"
  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "log"
  "os"
)

func ConnectDb() *gorm.DB {
  err := godotenv.Load()
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASS")
  dbHost := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")

  dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err.Error())
  }

  if err := db.AutoMigrate(&migrate.User{}, &migrate.Password{}); err != nil {
    return nil
  }

  return db
}