package leetcode

func IsMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	first := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (first && IsMatch(s[1:], p))
	} else {
		return first && IsMatch(s[1:], p[1:])
	}
}
