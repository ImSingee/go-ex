package glob

import "path/filepath"

// Match 执行通配符匹配
func Match(pattern, name string) bool {
	if result, err := filepath.Match(pattern, name); err == nil {
		return result
	} else {
		return pattern == name
	}
}

// MatchIn 返回给定列表中是否有满足 pattern 的元素
func MatchIn(pattern string, list []string) bool {
	for _, name := range list {
		if Match(pattern, name) {
			return true
		}
	}
	return false
}

// MatchPatterns 返回 patterns 中是否存在一个模式可以与 data 进行 Match
func MatchPatterns(data string, patterns []string) bool {
	for _, pattern := range patterns {
		if Match(pattern, data) {
			return true
		}
	}
	return false
}

// MatchListPatterns 对 list 中的每一个元素进行 MatchPatterns，有任一匹配则返回 true
func MatchListPatterns(list []string, patterns []string) bool {
	for _, data := range list {
		if MatchPatterns(data, patterns) {
			return true
		}
	}

	return false
}

// Find 返回给定列表中符合 pattern 的子列表
func Find(pattern string, list []string) []string {
	result := make([]string, 0, len(list))
	for _, name := range list {
		if Match(pattern, name) {
			result = append(result, name)
		}
	}
	return result
}
