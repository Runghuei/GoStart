package utest

import (
	"testing"
)

// go test 命令是一个按照一定约定和组织的测试代码驱动测试， 在包目录中 所有已_test.go结尾的文件都会被视为测试文件。
// 不会被go build 编译到可执行文件中
// 测试有四种类型： 1. 测试函数 2. 基准函数 3. 实例函数 4. 模糊测试
// 其他的测试框架， 测试期望， 更复杂的期望， 当我没去调用gorm进行查询的时候 我们知道结果是正确的呢？ 我先录制一些数据到表中， 查询出来对比一下就行了
// 测试如果底层依赖了其他的组件 库， 模糊？ mock，fake等
func TestName(t *testing.T) {
	t.Error("失败")
}
