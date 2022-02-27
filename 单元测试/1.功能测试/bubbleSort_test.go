package __功能测试

import (
	"reflect"
	"testing"
)

/*
测试语法
   	不需要main包，任何一个包，都可以给任何一个函数写测试函数

   	测试文件命名 : 必须以_test.go结尾
   	导入需要的包 : import testing
   	测试函数命名 : 必须以Test开头，后面的函数名是要测试的函数，而且不能以小写字母开头
	功能测试参数 : testing.T


如何执行测试函数
	如果用IDE，应该和有相应的功能,不过也有可能报错，原因也可能是没有相关依赖
	如果手动执行
		go test  				 会执行一个包下的所有测试文件，以及它的所有测试函数
		go test file_test.go 	 执行单个测试文件，会报错(测试函数没有定义)
								 报错原因：把测试文件当成了一个包(command-line-arguments),go test给定的
								 但是包中没有定义测试函数,所以报错了
								 执行命令时加入这个测试文件需要引用的源码文件
								 在命令行后方的文件都会被加载到command-line-arguments中进行编译
								 go test  file_test.go file.go

	go test -v 可以看具体案例的执行情况
*/

// BubbleSort函数的功能测试函数
func TestBubbleSort(t *testing.T) {
	// got 是调用测试函数
	got := []int{3, 6, 7, 0, 9, 1, 2, 5, 8, 4}
	BubbleSort(got)

	// want 是测试函数的执行结果
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 比较函数执行的结果是否和want相同
	// 由于切片不能直接比较，可以调用reflect.DeepEqual进行引用类型的比较

	// 如果执行结果不相同的处理
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v got:%v", want, got)
	}
}
