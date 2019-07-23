# String-Dispatch 

String-Dispatch是一个简单字符串分发组件，基于基数树实现字符串分发对应处理函数的分发器，类似于http router

代码基于[HttpRouter](https://github.com/julienschmidt/httprouter) 

## Usage
拷贝[tree](https://github.com/Duke-wei/string-dispatch/blob/master/tree/tree.go)代码，更换**Handle**为自定义接口，避免用例[example](https://github.com/Duke-wei/string-dispatch/blob/master/example/dispatch_test.go)中的接口类型断言

