package algo

// 后洗牌算法
import(
	"math/rand"
)


// 金额单位也是分
func AfterShuffle(count, amout int64) []int64 {
	inds := make([]int64, 0)

	remain := amout
	// 随机生成初级红包序列
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		inds = append(inds, x + min)
	}

	// 洗牌，打乱红包序列
	rand.Shuffle(len(inds), func(i, j int){
		inds[i], inds[j] = inds[j], inds[i]
	})
	return inds
}