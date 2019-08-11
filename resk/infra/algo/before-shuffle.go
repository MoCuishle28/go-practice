package algo

// 先洗牌算法
// 还是会先大后小
import(
	"time"
	"math/rand"
)


func BeforeShuffle(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算最大可调度金额
	max := amount - min*count
	// 生成红包种子金额序列
	seeds := make([]int64, 0)
	// 红包种子金额序列长度 = 3 ~ 1/2*count
	size := count/2
	if size < 3 {
		size = 3
	}
	for i := int64(0); i < size; i++ {
		x := max / (i + 1)
		seeds = append(seeds, x)
	}

	rand.Seed(time.Now().UnixNano())
	// 从种子金额序列中选出一个作为随机基数
	idx := rand.Int63n(int64(len(seeds)))
	// 使用随机基数最为最大数，随机出一个数字作为红包金额序列元素
	x := rand.Int63n(seeds[idx])
	return x + min
}