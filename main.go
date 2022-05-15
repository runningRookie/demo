package main

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
//	go func() {
//		_, _ = io.Copy(os.Stdout, conn)
//		log.Println("done")
//		done <- struct{}{}
//	}()
//	mustCopy(conn, os.Stdin)
//	conn.Close()
//	<-done
//}
//
//func mustCopy(dst io.Writer, src io.Reader) {
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

// 管道
//func main() {
//	naturals := make(chan int)
//	squares := make(chan int)
//	// counter
//	go func() {
//		for x := 0; x <= 10; x++ {
//			naturals <- x
//		}
//		close(naturals)
//	}()
//
//	// squarer
//	go func() {
//		for x := range naturals {
//			squares <- x * x
//		}
//		close(squares)
//	}()
//
//	// printer （在主 goroutine 中）
//	for x := range squares {
//		fmt.Println(x)
//	}
//}
