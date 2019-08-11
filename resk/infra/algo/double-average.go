package algo

// 二倍均值算法
// 红包金额分布概率在均值附近
import(
	"time"
	"math/rand"
)


func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算最大可用金额
	max := amount - min*count
	// 计算最大可用平均值
	avg := max/count
	// 二倍均值, 再加上最小金额（防止avg是0值）
	avg2 := avg*2 + min

	rand.Seed(time.Now().UnixNano())
	// 随机红包金额序列元素，把二倍均值作为随机的最大数
	x := rand.Int63n(avg2)
	return x
}