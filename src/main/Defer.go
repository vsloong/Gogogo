package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))

		//注意下面的写法，匿名函数
		defer func(n int) {
			fmt.Printf("%d ", n)
		}(fibonacci(i))
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

/*
defer语句仅能被放置在函数或方法中。

func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}


函数readFile的功能是读出指定文件或目录（以下统称为文件）本身的内容并将其返回，
同时当有错误发生时立即向调用方报告。
其中，os和ioutil（导入路径是io/ioutil）代表的都是Go语言标准库中的代码包。

请注意这个函数中的倒数第二条语句。
我们在打开指定文件且未发现有错误发生之后，紧跟了一条defer语句。
其中携带的表达式语句表示的是对被打开文件的关闭操作。
注意，当这条defer语句被执行的时候，其中的这条表达式语句并不会被立即执行。
它的确切的执行时机是在其所属的函数（这里是readFile）的执行即将结束的那个时刻。
也就是说，在readFile函数真正结束执行的前一刻，file.Close()才会被执行。
这也是defer语句被如此命名的原因。
我们在结合上下文之后就可以看出，

语句defer file.Close()的含义是在打开文件并读取其内容后及时地关闭它。
该语句可以保证在readFile函数将结果返回给调用方之前，
那个文件或目录一定会被关闭。

这实际上是一种非常便捷和有效的保险措施。
更为关键的是，无论readFile函数正常地返回了结果还是由于在其执行期间有运行时恐慌发生而被剥夺了流程控制权，
其中的file.Close()都会在该函数即将退出那一刻被执行。
这就更进一步地保证了资源的及时释放。

!!!!!
注意，当一个函数中存在多个defer语句时，
它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序。


 1. defer携带的表达式语句代表的是对某个函数或方法的调用。
这个调用可能会有参数传入，比如：fmt.Print(i + 1)。如果代表传入参数的是一个表达式，
那么在defer语句被执行的时候该表达式就会被求值了。
注意，这与被携带的表达式语句的执行时机是不同的。请揣测下面这段代码的执行：

func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}
    它在被执行之后，标准输出上打印出1 2 3 4 40 30 20 10 。

 2. 如果defer携带的表达式语句代表的是对匿名函数的调用，
那么我们就一定要非常警惕。请看下面的示例：

func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}
    deferIt4函数在被执行之后标出输出上会出现5555，而不是4321。
原因是defer语句携带的表达式语句中的那个匿名函数包含了对外部（确切地说，是该defer语句之外）
的变量的使用。注意，等到这个匿名函数要被执行（且会被执行4次）的时候，
包含该defer语句的那条for语句已经执行完毕了。此时的变量i的值已经变为了5。
因此该匿名函数中的打印函数只会打印出5。
正确的用法是：把要使用的外部变量作为参数传入到匿名函数中。修正后的deferIt4函数如下：

func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func(n int) {
            fmt.Print(n)
        }(i)
    }
}
 */
