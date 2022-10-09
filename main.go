package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	length             = 10000000
	ListRedisKey       = "LIST"
	SortedSetsRedisKey = "SORTED_SETS"
)

var (
	redisInit    bool
	redisInitStr string = "init"
)

func init() {
	flag.BoolVar(&redisInit, redisInitStr, false, "Redis initialize")
}

func newRedisConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func initDataset(conn redis.Conn) {
	// conn.Do("FLUSHALL")
	// for i := 0; i < length; i++ {
	// 	conn.Do("LPUSH", ListRedisKey, i)
	// 	conn.Do("ZADD", SortedSetsRedisKey, i, i)
	// }
}

func benchmark(conn redis.Conn, offset int) error {
	prefix := "benchmark"
	if err := os.Mkdir(prefix, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	listBench, err := os.OpenFile(fmt.Sprintf("./%v/list_%v", prefix, offset), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer listBench.Close()

	setBench, err := os.OpenFile(fmt.Sprintf("./%v/set_%v", prefix, offset), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer setBench.Close()

	for i := 0; i < length; i += offset {
		start := time.Now().UnixNano()
		conn.Do("LRANGE", ListRedisKey, i, i+offset-1)
		end := time.Now().UnixNano()

		listBench.WriteString(fmt.Sprintf("%v %v\n", i, end-start))

		start = time.Now().UnixNano()
		conn.Do("ZRANGE", SortedSetsRedisKey, i, i+offset-1)
		end = time.Now().UnixNano()
		setBench.WriteString(fmt.Sprintf("%v %v\n", i, end-start))
	}

	return nil
}

func main() {
	flag.Parse()

	conn, err := newRedisConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	flag.Visit(func(f *flag.Flag) {
		if f.Name == redisInitStr {
			initDataset(conn)
			fmt.Printf("Successfully initialize redis data of %v entries\n", length)
		}
	})

	offsets := []int{100, 1000, 10000}
	for _, offset := range offsets {
		if err := benchmark(conn, offset); err != nil {
			panic(err)
		}
		fmt.Printf("Successfully benchmark redis data of offset %v\n", offset)
	}
}
