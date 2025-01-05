package utils

// Diff 计算两个数组中的key变更元素
func Diff(cur, old []string) []string {
	set := make(map[string]bool)
	n := make([]string, 0)
	for _, v := range old {
		set[v] = true
	}

	for _, v := range cur {
		if !set[v] {
			n = append(n, v)
		}
	}
	return n
}
