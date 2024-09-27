package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type requestInfo struct {
	count    int
	timer    *time.Timer
}

// rateLimiter holds the request count for each IP.
type rateLimiter struct {
    visits map[string]*requestInfo
    lock   sync.Mutex
	limit  int
	window time.Duration
}

// newRateLimiter initializes a new rateLimiter.
func NewRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		visits: make(map[string]*requestInfo),
		limit:  limit,
		window: window,
	}
}
// limitMiddleware limits each IP to 3 requests to the endpoint.
func (rl *rateLimiter) RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.lock.Lock()
		info, exists := rl.visits[ip]
		if !exists {
			timer := time.AfterFunc(rl.window, func() {
				rl.lock.Lock()
				delete(rl.visits, ip)
				rl.lock.Unlock()
			})
			info = &requestInfo{
				count: 1,
				timer: timer,
			}
			rl.visits[ip] = info
		} else {
			info.count++
			info.timer.Reset(rl.window)
		}
		rl.lock.Unlock()

		if info.count > rl.limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}

		c.Next()
	}
}