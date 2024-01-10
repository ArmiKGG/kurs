package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func Top10(input string) []string {
	input = strings.TrimSpace(input)
	sliceOfStrings := strings.Split(input, " ")
	result := map[string]int{}
	top10Result := map[int][]string{}
	for _, s := range sliceOfStrings {
		_, ok := result[s]
		if ok {
			result[s] += 1
		} else {
			result[s] = 1
		}
	}

	var ss []kv
	for k, v := range result {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	top10 := []kv{}
	for _, kv := range ss {
		top10 = append(top10, kv)
		if len(top10) == 10 {
			break
		}
	}

	for _, kv := range top10 {
		top10Result[kv.Value] = append(top10Result[kv.Value], kv.Key)
	}

	for key, _ := range top10Result {
		sort.Strings(top10Result[key])
	}

	keys := make([]int, 0, len(top10Result))
	for k := range top10Result {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	final := []string{}
	for _, key := range keys {
		final = append(final, top10Result[key]...)
	}

	return final
}
