package v3

import (
	"errors"
	"testing"
	"time"
)

func TestV3Log1(t *testing.T) {
	Infof("hello info log:info=%v", "v3 info logger")
	Errorf("hello error log:error=%v", "v3 err log")
}

func TestV3Log2(t *testing.T) {
	EnableFileLog(true)
	SetRotateType(RotateMinute)
	SetFileLog("../logging.log")

	ticker := time.NewTicker(time.Millisecond * 50)
	for {
		select {
		case <-ticker.C:
			InfoKV(KV{"time": time.Now().Unix()}, "InfoKV")
			Infof("hello infof:msg=%v", "nice infof")
			Errorf("hello errorf:err=%v", errors.New("bad request"))
			ErrorKV(KV{"error": errors.New("bad message for you").Error()}, "ErrorKV")
		}
	}

}
