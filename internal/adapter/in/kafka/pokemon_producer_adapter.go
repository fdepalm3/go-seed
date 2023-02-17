package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main(){
	conn ,_ :=kafka.DialLeader (context.Background(),"tcp", "127.0.0.1:9092", "kafka-topic",0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: []byte("Pokemon attack ")})

}