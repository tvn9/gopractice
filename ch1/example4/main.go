// Getting and setting environment variables with default values
package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	// Set the environment variable
	os.Setenv(key, "postgres://as:tv.nguyen9@icloud.com/pg?sslmode=verify-full")
	val := GetEnvDefault(key, "postgres://as:tv.nguyen9@localhost/pg?sslmode=verify-full")

	log.Println("The value is :" + val)

	os.Unsetenv(key)
	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")
	log.Println("The default value is :" + val)
}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
