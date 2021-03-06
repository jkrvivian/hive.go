package async

import (
	"crypto/rand"
	"golang.org/x/crypto/blake2b"
	"sync"
	"testing"
)

func cpuHeavyFunc() {
	randomData := make([]byte, 102400, 102400)
	_, _ = rand.Read(randomData)

	blake2b.Sum512(randomData)
}

func BenchmarkWorkerPool_goroutine(b *testing.B) {
	b.ReportAllocs()

	var wg sync.WaitGroup

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			cpuHeavyFunc()

			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkWorkerPool_Submit(b *testing.B) {
	b.ReportAllocs()

	var wp WorkerPool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wp.Submit(cpuHeavyFunc)
	}

	wp.ShutdownGracefully()
}
