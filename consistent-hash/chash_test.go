package consistent_hash

import (
	"log"
	"testing"
)

func TestChash1(t *testing.T) {
	hashRing := NewHashRing(100)
	hashRing.AddNode("localhost:10000", 1)
	hashRing.AddNode("localhost:10001", 2)
	hashRing.AddNode("localhost:10002", 1)
	hashRing.AddNode("localhost:10003", 2)
	hashRing.AddNode("localhost:10004", 3)

	nodeName := hashRing.GetNode("/api/hello")
	log.Printf("node name:%v", nodeName)

	nodeName = hashRing.GetNode("/api/ping")
	log.Printf("node name:%v", nodeName)

	nodeName = hashRing.GetNode("/api/book/1")
	log.Printf("node name:%v", nodeName)

	nodeName = hashRing.GetNode("/api/user/info")
	log.Printf("node name:%v", nodeName)

}
