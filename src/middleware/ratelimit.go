package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

var (
	vistors    = make(map[string]*rate.Limiter)
	mu         sync.Mutex
	blockedIPs = make(map[string]time.Time)
)

func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()

		if until, blocked := blockedIPs[ip]; blocked && time.Now().Before(until) {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "IP temporarily blocked"})
			return
		}

		mu.Unlock()

		limiter := getVistor(c)

		if !limiter.Allow() {
			mu.Lock()

			blockedIPs[ip] = time.Now().Add(10 * time.Second)

			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, blocked for 10s"})
			return
		}

		c.Next()
	}
}

func getVistor(c *gin.Context) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	ip := c.ClientIP()

	limiter, exists := vistors[ip]

	if !exists {
		limiter = rate.NewLimiter(1, 5)
		vistors[ip] = limiter
	}

	return limiter
}
