package exstrings

func InStringList(list []string, want string) bool {
	for _, v := range list {
		if v == want {
			return true
		}
	}

	return false
}

// PopInStringList 如果元素在列表中，则返回 true 与删除元素后的列表，否则返回 false 与原列表
func PopInStringList(list []string, want string) (bool, []string) {
	for i, v := range list {
		if v == want {
			return true, append(list[:i], list[i+1:]...)
		}
	}
	return false, list
}

// CountNonEmpty 返回非空字符串的个数
func CountNonEmpty(ss ...string) (c int) {
	for _, s := range ss {
		if s != "" {
			c++
		}
	}

	return
}
