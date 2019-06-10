# goablite
Golang语言实现的HTTP压测工具，类似ApacheBench

# 1. 单机单用户版--alone.go

参考：https://www.jianshu.com/p/7f28aad6ecc5

介绍：输入2个参数：并发数 目标url

使用方法：在对应目录下，使用go build进行编译，然后运行alone.exe n url，不限连接次数，需手动停止。

# 2. 服务端与客户端分离版

参考：http://www.voidcn.com/article/p-hxfyoovt-bpb.html
