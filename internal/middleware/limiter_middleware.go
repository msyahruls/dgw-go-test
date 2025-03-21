package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	mMemory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  60, // max 60 requests/minute
	}
	store := mMemory.NewStore()
	middleware := mGin.NewMiddleware(limiter.New(store, rate))

	return middleware
}
