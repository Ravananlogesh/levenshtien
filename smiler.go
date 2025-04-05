package levenshtien

import "math"

// ============ STRING MATCH ===============

func Min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

/* Ultra-fast Levenshtein using single slice to reduce the memory footprint
and improve cache performance. This is a simplified version of the
algorithm that only works for ASCII characters. It is not as fast as
the full version, but it is much simpler and easier to understand.
It is also not as fast as the full version, but it is much simpler and
easier to understand. It is also not as fast as the full version, but
it is much simpler and easier to understand. It is also not as fast as */
func LevenshteinDistance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}
	if lenA > lenB {
		a, b = b, a
		lenA, lenB = lenB, lenA
	}
	row := make([]int, lenA+1)
	for i := 0; i <= lenA; i++ {
		row[i] = i
	}
	for j := 1; j <= lenB; j++ {
		prev := row[0]
		row[0] = j
		for i := 1; i <= lenA; i++ {
			old := row[i]
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			ins := row[i-1] + 1
			del := row[i] + 1
			sub := prev + cost
			min := ins
			if del < min {
				min = del
			}
			if sub < min {
				min = sub
			}
			row[i] = min
			prev = old
		}
	}
	return row[lenA]
}
// ============ STRING SIMILARITY ===============
// Similarity returns a value between 0.0 and 100.0

func Similarity(a, b string) float64 {
	dist := LevenshteinDistance(a, b)
	maxLen := max(len(b), len(a))
	if maxLen == 0 {
		return 100.0
	}
	return (1.0 - float64(dist)/float64(maxLen)) * 100.0
}

// ============ NUMBER MATCH ===============
// MatchNumber returns a value between 0.0 and 100.0
// It uses the absolute difference between the two numbers
// and the maximum of the two numbers to calculate the similarity score.
// If the two numbers are equal, it returns 100.0.
// If both numbers are 0, it returns 0.0.
// If one of the numbers is 0, it returns 0.0.
// The similarity score is calculated as follows:
func MatchNumber(a, b float64) float64 {
	if a == b {
		return 100.0
	}
	diff := math.Abs(a - b)
	max := math.Max(math.Abs(a), math.Abs(b))
	if max == 0 {
		return 0.0
	}
	return (1.0 - diff/max) * 100.0
}

// ============ LIST MATCH (STRING) ===============

type MatchResult struct {
	String     string
	Similarity float64
}
// MatchStringList returns the best match from a list of strings
// and the similarity score.
// It uses the Levenshtein distance to calculate the similarity score.
// The best match is the one with the highest similarity score.
// If there are multiple matches with the same score, the first one is returned.
// The input string is compared to each string in the list and the
// similarity score is calculated.
// The function returns the best match and the similarity score.

func MatchStringList(input string, list []string) MatchResult {
	best := MatchResult{"", -1}
	for _, s := range list {
		score := Similarity(input, s)
		if score > best.Similarity {
			best = MatchResult{s, score}
		}
	}
	return best
}

// ============ LIST MATCH (NUMBER) ===============

type MatchResultNum struct {
	Number     float64
	Similarity float64
}
// MatchNumberList returns the best match from a list of numbers
// and the similarity score.
// It uses the absolute difference between the two numbers
// and the maximum of the two numbers to calculate the similarity score.
// The best match is the one with the highest similarity score.
// If there are multiple matches with the same score, the first one is returned.
// The input number is compared to each number in the list and the
// similarity score is calculated.
func MatchNumberList(input float64, list []float64) MatchResultNum {
	best := MatchResultNum{0, -1}
	for _, n := range list {
		score := MatchNumber(input, n)
		if score > best.Similarity {
			best = MatchResultNum{n, score}
		}
	}
	return best
}

// ============ HYBRID MATCH ===============
// You can give 0.0 to 1.0 weight for string vs number match
func CombinedMatch(s1, s2 string, n1, n2 float64, weight float64) float64 {
	strSim := Similarity(s1, s2)
	numSim := MatchNumber(n1, n2)
	return strSim*weight + numSim*(1.0-weight)
}
