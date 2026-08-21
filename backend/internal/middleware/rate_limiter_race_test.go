package middleware

import (
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRateLimiterConcurrentSnapshotAndLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rl := NewRateLimiter(100000)

	const workers = 8
	const loops = 200
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			<-start
			for j := 0; j < loops; j++ {
				switch worker % 6 {
				case 0:
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Request = httptest.NewRequest("GET", "/", nil)
					c.Request.RemoteAddr = "10.0.0.1:80"
					rl.Limit()(c)
				case 1:
					rl.Snapshot()
				case 2:
					rl.Count()
				case 3:
					rl.Reset("10.0.0.1")
				case 4:
					rl.Peek("10.0.0.1")
				case 5:
					rl.ClearAll()
				}
			}
		}(i)
	}
	close(start)
	wg.Wait()
}
