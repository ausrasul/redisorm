# Redis wrapper for Golang

### Installation
```
go get github.com/ausrasul/redisorm
```

### Usage:

Use this library to directly save structs to Redis and get structs from Redis.
Connection pool management is built in.

### Usage example:

```
package main

import (
   "github.com/ausrasul/redisorm"
)

type User struct {
	Name string
	Id   int
}

user := &User{
	Name: "test",
	Id: 1,
}
func main(){
	// Configure the package

	redisorm.Configure(
		map[string]interface{}{
			"poolMaxIdle":   10,
			"poolMaxActive": 60,
			"port":          "6379",
			"db":             1, // optional, default 0
		},
	)

	err := redisorm.Set(user.Name, user)  // user.Name is just a key
	err := redisorm.Get(user.Name, user)
	//handle errors.
	return err
}

```
