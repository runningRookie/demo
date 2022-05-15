package main

import (
	"fmt"
	"time"
)

// 通道
// 创建 make 无缓冲通道、有缓冲通道
// 使用 read/write <- -> 阻塞
// 销毁 close
//func main() {
//	nums := make(chan int, 10)
//	nums <- 10
//	go func() {
//		fmt.Println(<-nums)
//		nums <- 11
//	}()
//	time.Sleep(1 * time.Second)
//	fmt.Println(<-nums)
//}

//func main() {
//	conn, err := net.Dial("tcp", "localhost:8000")
//	if err != nil {
//		log.Fatal(err)
//	}
//	done := make(chan struct{})
//	log.Println("连接创建成功")
//	go func() {
//		// 从连接读取数据到标准输出
//		_, err = io.Copy(os.Stdout, conn)
//		if err != nil {
//			log.Println(err)
//		}
//		log.Println("done")
//		done <- struct{}{}
//	}()
//	// 从输入copy到连接
//	mustCopy(conn, os.Stdin)
//	// 连接关闭
//	log.Println("关闭连接，写半边连接")
//	err = conn.(*net.TCPConn).CloseWrite()
//	if err != nil {
//		log.Fatal(err)
//	}
//	<-done
//}

//func mustCopy(dst io.Writer, src io.Reader) {
//	log.Println("copy 标准输入到连接")
//	if _, err := io.Copy(dst, src); err != nil {
//		log.Fatal(err)
//	}
//}

// 通道chan
// 管道是不一样的操作

//服务器回声
//func main() {
//	listen, err := net.Listen("tcp", "localhost:8000")
//	if err != nil {
//		log.Fatal(err) // server 创建失败
//	}
//	for { // 不停的接受用户请求
//		con, err := listen.Accept()
//		if err != nil {
//			log.Println(err) // 例如，连接终止
//			continue
//		}
//		go handleConn(con) // 一次处理一个连接
//	}
//}
//
//func handleConn(con net.Conn) {
//	defer con.Close()
//	for {
//		// 写入con
//		_, err := io.WriteString(con, fmt.Sprintf("%s\n", time.Now().Format("15:04:05")))
//		if err != nil {
//			return // 例如：连接断开
//		}
//		time.Sleep(time.Second * 1)
//	}
//}

// 并发问题
//
//func main() {
//	go remind(200 * time.Millisecond)
//	const n = 45
//	fibonacci := f(n)
//	fmt.Printf("\r fibonacci(%d)=%d\n", n, fibonacci)
//}
//
//func f(n int) int {
//	if n < 2 {
//		return n
//	}
//	return f(n-2) + f(n-1)
//}
//
//func remind(duration time.Duration) {
//	for {
//		for _, r := range `-\|/` {
//			fmt.Printf("\r%c", r)
//			time.Sleep(duration)
//		}
//	}
//}

//管道
//func main() {
//	naturals := make(chan int)
//	squares := make(chan int)
//	// counter
//	go func() {
//		for x := 0; x < 100; x++ {
//			naturals <- x
//		}
//		close(naturals)
//	}()
//
//	// squarer
//	go func() {
//		// 有更加通用的写法，rang
//		for {
//			x, ok := <-naturals
//			if !ok {
//				break // 关闭通道并且已经读完
//			}
//			squares <- x * x
//		}
//		close(squares)
//	}()
//
//	// printer （在主 goroutine 中）
//	for x := range squares {
//		fmt.Println(x)
//		time.Sleep(100 * time.Millisecond)
//	}
//}

// 关闭通道作为一个广播机制？
//func main() {
//	sig := make(chan struct{})
//
//	go func() {
//		<-sig
//		fmt.Println("我是一号")
//	}()
//
//	go func() {
//		<-sig
//		fmt.Println("我是二号")
//	}()
//	go func() {
//		<-sig
//		fmt.Println("我是三号")
//	}()
//
//	for i := 0; i < 10; i++ {
//		time.Sleep(1 * time.Second)
//	}
//	close(sig)
//	time.Sleep(1 * time.Second)
//}

//func main() {
//	ch := make(chan int, 1) // 如果没有容量非常容易发生死锁
//	for i := 0; i < 10; i++ {
//		select {
//		case x := <-ch:
//			fmt.Println(x)
//		case ch <- i:
//		}
//	}
//	log.Println("finished!")
//}

func main() {
	//tick := time.Tick(1 * time.Second)
	//for range tick {
	//	fmt.Println(time.Now().Format("15:04:05"))
	//}

	ticker := time.NewTicker(1 * time.Second)
	count := 0
	for range ticker.C {
		fmt.Println(time.Now().Format("15:04:05"))
		count++
		if count > 10 {
			break
		}
	}
	ticker.Stop()
}
