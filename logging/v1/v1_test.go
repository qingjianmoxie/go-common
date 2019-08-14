package v1

import "testing"

func TestV1Log1(t *testing.T) {
	PrintKV(KV{
		"a": "hello",
		"b": 123}, "haha")
}

func TestV1Log2(t *testing.T)  {
	EnableLogFile(true)
	SetLogfile("logging.log")

	Infof("This is a Info msg:%v", "haha")
}

func TestMain(m *testing.M)  {
	m.Run()
}
