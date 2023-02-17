package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "kafka-topic", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	batchMessage := conn.ReadBatch(1e3, 1e9) // 1e3 1000

	bytes := make([]byte, 1e3)

	for {
		_, err := batchMessage.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}

}
