package protocol

import (
	"testing"	
	"github.com/goodkele/mgnet/library/module/types"
	"sync"
	//"runtime"
	//"time"
)

type epool struct {
	epoolChan	chan []byte
}

func NewEpool() (*epool) {
	e := &epool {
		epoolChan	:	make(chan[]byte, 1024),
	}
	return e
}

func (this *epool) Read(p []byte) (n int, err error) {
	c := <- this.epoolChan
	copy(p, c)
	return len(p), nil
}

func (this *epool) Write(p []byte) (n int, err error) {
	this.epoolChan <- p
	return len(p), nil
}

var (
	e = NewEpool()
)

func Test_Protocol(t *testing.T) {
	codecType := &CodecType{}
	encode := codecType.NewEncoder(e)
	decode := codecType.NewDecoder(e)

	var waitSync sync.WaitGroup
	waitSync.Add(100)

	go func() {
		for i:=0; i<100; i++ {
			if i % 2 == 0 {
				seReq := &types.SearchRequest{"name1", int32(i)+1, int32(i)+2}
				encode.Encode(seReq)
			} else {
				routing := &types.Routing{1,2,3,"4",[]byte("5"),6,7,"error"}	
				encode.Encode(routing)
			}
			// time.Sleep(100 * time.Millisecond)
		}
	}()
	
	//runtime.Gosched()
	
	go func() {
		var i int
		for {
			if i % 2 == 0 {
				seReq := &types.SearchRequest{}
				decode.Decode(seReq)
				//fmt.Println(seReq)
			} else {
				routing := &types.Routing{}	
				decode.Decode(routing)
				//fmt.Println(routing)
			}
			i++
			waitSync.Done()
		}
	}()

	waitSync.Wait()	
}

func Benchmark_Protocol(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	codecType := &CodecType{}
	encode := codecType.NewEncoder(e)
	decode := codecType.NewDecoder(e)

	var waitSync sync.WaitGroup
	var n int
	n = 1000000
	
	waitSync.Add(n)
	go func() {
		for i:=0; i<n; i++ {
			if i % 2 == 0 {
				seReq := &types.SearchRequest{"name1", int32(i)+1, int32(i)+2}
				encode.Encode(seReq)
			} else {
				routing := &types.Routing{1,2,3,"4",[]byte("5"),6,7,"error"}	
				encode.Encode(routing)
			}
		}
	}()
	
	//runtime.Gosched()
	
	go func() {
		var i int
		for {
			if i % 2 == 0 {
				seReq := &types.SearchRequest{}
				decode.Decode(seReq)
			} else {
				routing := &types.Routing{}	
				decode.Decode(routing)
			}
			
			i++
			
			waitSync.Done()
		}
	}()
	
	waitSync.Wait()	
}
