**通道**
1. 创建 make 无缓冲通道、有缓冲通道
2. 使用 read/write <- -> 阻塞
3. 销毁 close

常用工具方法
* `io.Copy`
* `io.WriteString`
* `os.Stdin`
* `os.Stdout`
* `os.StdErr`

常见接口
* io.Writer
* io.Reader

**练习 8.3**

```go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	log.Println("连接创建成功")
	go func() {
		// 从连接读取数据到标准输出
		_, err = io.Copy(os.Stdout, conn)
		if err != nil {
			log.Println(err)
		}
		log.Println("done")
		done <- struct{}{}
	}()
	// 从输入copy到连接
	mustCopy(conn, os.Stdin)
	// 连接关闭
	log.Println("关闭连接，写半边连接")
	err = conn.(*net.TCPConn).CloseWrite()
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	log.Println("copy 标准输入到连接")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```

