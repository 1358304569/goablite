## goablite

Golang语言实现的缩减版HTTP压测工具，类似ApacheBench

名字很简单，go语言编写，类似ab，用的LiteIDE


**参数：**
	-n:连接数
	-c:并发数
	-p:CPU数
	-u:URL
**输出：**
	Complete requests:成功建立连接的请求数
	Failed requests:连接失败的请求数
	UseTime:一共耗费的时间（s）
	RPS:每秒完成的HTTP连接数（request per second）
