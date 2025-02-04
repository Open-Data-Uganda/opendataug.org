package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"opendataug.org/errors"
)

type RateLimiter struct {
	ips    map[string]*rate.Limiter
	mu     sync.RWMutex
	rate   rate.Limit
	burst  int
	expiry time.Duration
}

func NewRateLimiter(r rate.Limit, burst int, expiry time.Duration) *RateLimiter {
	return &RateLimiter{
		ips:    make(map[string]*rate.Limiter),
		rate:   r,
		burst:  burst,
		expiry: expiry,
	}
}

func (rl *RateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.ips[ip] = limiter
	}

	return limiter
}

func RateLimit(requests int, per time.Duration, burst int) gin.HandlerFunc {
	rateLimiter := NewRateLimiter(rate.Every(per/time.Duration(requests)), burst, 1*time.Hour)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rateLimiter.getLimiter(ip)

		if !limiter.Allow() {
			err := errors.NewRateLimitError("Rate limit exceeded. Please try again later.")
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
