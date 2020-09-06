# Redis wrapper for Golang

### Installation
```
go get github.com/ausrasul/redisorm
```

### Usage:

Use this library to directly save structs to Redis and get, delete structs from Redis.
Connection pool management is built in.

### Usage example:

```
package main

import (
	"github.com/ausrasul/redisorm"
)

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
	user := struct{
		Name string
		Id int
	}{
		Name: "test",
		Id: 1,
	}
	err := redisorm.Set(user.Name, user)  // user.Name is the key
	err := redisorm.Get(user.Name, &user)
	// err != nil means unmarshaling error, db error
	// err.Error() == "no_data" means no data.
	cnt, err := redisorm.Del(user.Name)
	// cnt number of items deleted, err != nil if db errors.
	cnt, err = redisorm.Del(["key1", "key2", "key3"])
	// can delete multiple keys.
	return err
}

```
