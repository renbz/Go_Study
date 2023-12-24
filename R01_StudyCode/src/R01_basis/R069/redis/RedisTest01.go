package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.DiaL("tcp", "82.157.50.241:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer conn.close()
	fmt.Println("conn suc...", conn)
}
