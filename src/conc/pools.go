package conc

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

/*
when working with a Pool, just remember the following points:
	• When instantiating sync.Pool, give it a New member variable that is thread-safe
	when called.
	• When you receive an instance from Get, make no assumptions regarding the
	state of the object you receive back.
	• Make sure to call Put when you’re finished with the object you pulled out of the
	pool. Otherwise, the Pool is useless. Usually this is done with defer.
	• Objects in the pool must be roughly uniform in makeup.
*/

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}
