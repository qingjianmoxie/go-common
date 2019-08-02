package snowflake

import (
	"log"
	"testing"
)

func TestSnowflake1(t *testing.T) {
	node, err := NewNode(100)
	if err != nil {
		log.Printf("fail to new node:%v", err)
		return
	}

	for i := 1; i < 100; i++ {
		log.Printf("get id:%v", node.Generate().Int64())
	}

}
