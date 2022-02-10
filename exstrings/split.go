package exstrings

import "strings"

func DeepSplit(s string) []string {
	result := make([]string, 0, 5)

	splits := strings.Split(s, " ")
	for _, split := range splits {
		split = strings.TrimSpace(split)

		if split == "" {
			continue
		}

		newSplits := strings.Split(s, ",")
		for _, newSplit := range newSplits {
			newSplit = strings.TrimSpace(newSplit)

			if newSplit == "" {
				continue
			}

			result = append(result, newSplit)
		}
	}

	return result
}

func DeepSplits(ss []string) []string {
	result := make([]string, 0, len(ss)<<2)

	for _, s := range ss {
		result = append(result, DeepSplit(s)...)
	}

	return result
}
