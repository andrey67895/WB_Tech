package grep

func Search(lines []string, m *Matcher, opt Options) []Result {
	matches := make([]bool, len(lines))
	for i, l := range lines {
		matches[i] = m.Match(l)
	}

	if opt.Count {
		count := 0
		for _, ok := range matches {
			if ok {
				count++
			}
		}
		res := make([]Result, count)
		return res
	}

	printed := make([]bool, len(lines))
	var res []Result

	for i, ok := range matches {
		if ok {
			start := i - opt.Before
			if start < 0 {
				start = 0
			}
			end := i + opt.After
			if end >= len(lines) {
				end = len(lines) - 1
			}
			for j := start; j <= end; j++ {
				if !printed[j] {
					res = append(res, Result{
						Line:    lines[j],
						LineNum: j + 1,
						Matched: matches[j],
					})
					printed[j] = true
				}
			}
		}
	}
	return res
}
