package algo

// 简单随机算法
// 缺点 容易产生很大的值，且前面领取的大红包概率较大
import(
	"time"
	"math/rand"
)


// 1分为最小金额
const min = int64(1)

// 算法有问题，会产生未到最后一个红包就减到0的情况
// 红包数量、红包金额(金额以分为单位， 如：1元 = 100分)
func SimpleRand(count, amount int64) int64 {
	// 红包剩余一个时
	if count == 1 {
		return amount
	}
	// 计算最大可调度金额
	max := amount - min*count
	// 设置随机种子
	rand.Seed(time.Now().Unix())
	// min~max 的随机值
	x := rand.Int63n(max) + min
	return x
}