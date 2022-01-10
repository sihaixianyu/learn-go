package goroutine

import (
	"testing"
)

func Test_echoReq(t *testing.T) {
	// 放到 main 函数调用 go test 里无法获取标准输入
	echoReq()
}
