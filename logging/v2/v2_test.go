package v2

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

type MStruct struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func (ms MStruct) String() string {
	bytes, err := json.Marshal(ms)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func TestStdLogger(t *testing.T) {
	log.Printf("Hello world")
	log.Printf("Hello %v", "Golang")
	m := MStruct{A: "Get", B: 123}
	log.Println("a", 1, m.String())
}

func TestStd2Logger(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Printf("Hello %v", "Logger")
	logger.Panicf("panic info")
}

func TestFileLogger(t *testing.T) {
	file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	logger.Printf("Hello File Log")
	logger.Printf("Just A OK:a=%v,b=%v", "123", 456)
}

func TestV2Log(t *testing.T) {
	Infof("Hello s=%v", "Golang")
	Errorf("Some Error Emerge:a=%v", 123)

	InfoKV(KV{"age": 1, "name": "jack"}, "Person Info")
	ErrorKV(KV{"uid": "u12344", "height": 177}, "Error Info")
}

func TestV2Log1(t *testing.T) {
	EnableFileLog(true)
	SetFileLog("../../logging.log")

	Infof("Hello s=%v", "Golang")
	Errorf("Some Error emerge:a=%v", 123)

	InfoKV(KV{"age": 1, "name": "jack"}, "Person Info")
	ErrorKV(KV{"uid": "u12344", "height": 177}, "Error Info")
}
