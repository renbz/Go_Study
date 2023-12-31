package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "82.157.50.241:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	fmt.Println("conn suc...", conn)
}
