# go 的垃圾回收机制

1. 引用计数
2. 标记清扫
3. 节点复制
4. 分代回收



golang 局部变量的生命周期，是从它们被声明开始，到不再继续被引用为止。如果函数返回指针，则内存逃逸，不被回收！！！
