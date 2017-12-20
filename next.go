package kmp

// 下标从0开始
// 当小标为j时与主串失配，
// 则主串当前的字符应该与k=next[j]进行比较,
// k为p[:j]（前面长度为j的字符串），相同前后缀最大长度(next[1]=0)
// next[0]=-1表示主串下标应该向前移一个位置
//
// h=next[j-1] 表示p[:j-1]相同前后缀最大长度
// 若p[h] == p[j-1] 则 k=next[j]=h+1
// 若p[h] != p[j-1] 则 判断 g = next[h]，g不小0;p[g] 是否等于 p[j-1]
func GetNext(p []byte) []int {
	cnt := len(p)
	if cnt == 0 {
		return nil
	}
	next := make([]int,cnt)
	next[0] = -1
	if cnt > 1 {
		next[1] = 0
	}
	for j := 2;j < cnt; j++ {
		h := next[j-1]
		for h >= 0 {
			if p[h] == p[j - 1] {
				next[j] = h + 1
			} else {
				h = next[h]
			}
		}
		if h == -1 {
			next[j] = 0
		}
	}
	return next
}
