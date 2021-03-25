package epolling

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"testing"
	"time"

	"github.com/mailru/easygo/netpoll"
)

func BenchmarkEpoll(b *testing.B) {
	poller, _ := netpoll.New(nil) // new go routine for waiting
	// var x int = 0
	lis, _ := net.Listen("tcp", ":1234")

	go func() {
		for {
			conn, _ := lis.Accept()
			desc := netpoll.Must(netpoll.HandleRead(conn))
			now := time.Now()

			poller.Start(desc, func(evet netpoll.Event) {
				if evet&netpoll.EventReadHup != 0 {
					poller.Stop(desc)
					conn.Close()
					return
				}
				go func() {

					data := make([]byte, 100)
					conn.Read(data)
					fmt.Printf("msg: %s, fd: %+d\n", string(data), now.UnixNano())
				}()
			})
		}

	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 50; i++ {
		conn2, _ := net.Dial("tcp", ":1234")
		io.WriteString(conn2, "test1")
		if i == 0 {
			go func() {
				for i := 0; i < 10; i++ {
					io.WriteString(conn2, "test2")
				}
			}()
		}

	}
	fmt.Printf("\nnum of threads: %d\n", runtime.NumGoroutine())
}

func BenchmarkBlock(b *testing.B) {
	lis, _ := net.Listen("tcp", ":2341")

	go func() {
		for {
			conn, _ := lis.Accept()

			go func(conn net.Conn) {
				// _ = make([]int64, 1000000)
				data := make([]byte, 100000)

				for {
					conn.Read(data)
				}
			}(conn)
		}
	}()

	time.Sleep(time.Millisecond * 1000)

	for i := 0; i < 50; i++ {
		conn2, _ := net.Dial("tcp", ":2341")
		io.WriteString(conn2, "test1")
	}
	fmt.Printf("\nnum of: %d\n", runtime.NumGoroutine())
}
