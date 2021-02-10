package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type BucketRule struct {

	// Key 自定义键值对名称
	Key string
	// FillInterval 间隔多久放N个令牌
	FillInterval time.Duration
	// Capacity 令牌桶的容量
	Capacity int64
	// Quantum 每次达到间隔时间后所放令牌具体数量
	Quantum int64
}

// LimiterIF 用于不同接口的限流器
type LimiterIF interface {
	// Key 获取对应限流器键值对名称
	Key(c *gin.Context) string
	// GetBucket 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// AddBuckets 新增多个令牌桶
	AddBuckets(rules ...BucketRule) LimiterIF
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}
