package algo

import(
	"testing"
)


func TestSimple(t *testing.T) {
	tests := [] struct {count, amount, remain int64} {
		{10, 100*100, 100*100},
		{12, 128*100, 128*100},
		{32, 621*100, 621*100},
		{5, 4*100, 4*100},
	}

	var sum int64 = 0
	for _, tt := range tests {
		for i := int64(0); i < tt.count; i++ {
			x := SimpleRand(tt.count-i, tt.amount)
			sum += x
			tt.amount = tt.amount - x
		}
		if tt.remain != sum {
			t.Errorf("test:%d, %d; got:%d, expected:%d\n", tt.count, tt.remain, sum, tt.remain)
		}
		sum = 0
	}
}


func TestDoubleAverage(t *testing.T) {
	tests := [] struct {count, amount, remain int64} {
		{10, 100*100, 100*100},
		{12, 128*100, 128*100},
		{32, 621*100, 621*100},
		{5, 4*100, 4*100},
	}

	var sum int64 = 0
	for _, tt := range tests {
		for i := int64(0); i < tt.count; i++ {
			x := DoubleAverage(tt.count-i, tt.amount)
			sum += x
			tt.amount = tt.amount - x
		}
		if tt.remain != sum {
			t.Errorf("test:%d, %d; got:%d, expected:%d\n", tt.count, tt.remain, sum, tt.remain)
		}
		sum = 0
	}
}


// 出错
// func TestBeforeShuffle(t *testing.T) {
// 	tests := [] struct {count, amount, remain int64} {
// 		{10, 100*100, 100*100},
// 		{12, 128*100, 128*100},
// 		{32, 621*100, 621*100},
// 		{5, 4*100, 4*100},
// 	}

// 	var sum int64 = 0
// 	for _, tt := range tests {
// 		for i := int64(0); i < tt.count; i++ {
// 			x := BeforeShuffle(tt.count-i, tt.amount)
// 			sum += x
// 			tt.amount = tt.amount - x
// 		}
// 		if tt.remain != sum {
// 			t.Errorf("test:%d, %d; got:%d, expected:%d\n", tt.count, tt.remain, sum, tt.remain)
// 		}
// 		sum = 0
// 	}
// }


// 出错
func TestAfterShuffle(t *testing.T) {
	tests := [] struct {count, amount int64} {
		{10, 100*100},
		{12, 128*100},
		{32, 621*100},
		{5, 4*100},
	}

	var sum int64 = 0
	for _, tt := range tests {
		arr := AfterShuffle(tt.count, tt.amount)
		for _, v := range arr {
			sum += v
		}
		if tt.amount != sum {
			t.Errorf("test:%+v; got:%d, arr:%v\n", tt, sum, arr)
		}
		sum = 0
	}
}