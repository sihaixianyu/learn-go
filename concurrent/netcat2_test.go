package concurrent

import (
	"testing"
)

func Test_call2(t *testing.T) {
	// 放到 main 函数调用 go test 里无法获取标准输入
	Call2()
}
