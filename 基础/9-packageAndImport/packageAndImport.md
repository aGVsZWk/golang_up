跟 python 一样啊，就是要在每个文件开始前拿 package 声明自己被其它文件导入的名字。

golang 有 gopath的，导包的时候，最好写绝对路径！！！不要写相对路径

每个文件夹（目录）用 package 相同名字，函数只有首字母大写才能被别的引用。

设置 GOBIN，然后用 go install 安装
