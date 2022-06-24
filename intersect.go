package glob_intersection

import (
	"fmt"
)

var Debug = false
var Col = 15

func Intersect(l, r string) bool {
	if l == r {
		return true
	}

	var reason string
	if Debug {
		fmt.Println("----------------------------------------------------------------------")
		reason = `initial`
	}

	intersected := _intersect([]rune(l), []rune(r), reason)

	if Debug {
		reason := `not intersected`
		if intersected {
			reason = `intersected`
		}
		fmt.Printf("|%*s| |%*s| %s\n", Col, ``, Col, ``, reason)
	}

	return intersected
}

func _intersect(l, r []rune, reason string) bool {
	if Debug {
		fmt.Printf("|%*s| |%*s| %s\n", Col, string(l), Col, string(r), reason)
	}

	if len(l) == 0 {
		if len(r) == 0 {
			return true
		}
		if r[0] != '*' {
			return false
		}
		if Debug {
			reason = fmt.Sprintf("expand right: * → %s", `∅`)
		}
		return _intersect(l, r[1:], reason)
	}

	if l[0] != '*' {
		if len(r) == 0 {
			return false
		}
		if l[0] == r[0] {
			if Debug {
				reason = fmt.Sprintf("advance both: %c ↔ %c", l[0], r[0])
			}
			return _intersect(l[1:], r[1:], reason)
		}
		if r[0] != '*' {
			return false
		}
		for i := 0; i <= len(l); i++ {
			if Debug {
				ls := `∅`
				if i > 0 {
					ls = string(l[:i])
				}
				reason = fmt.Sprintf("expand right: * → %s", ls)
			}
			if _intersect(l[i:], r[1:], reason) {
				return true
			}
		}
		return false
	}

	for j := 0; j <= len(r); j++ {
		if Debug {
			rs := `∅`
			if j > 0 {
				rs = string(r[:j])
			}
			reason = fmt.Sprintf("expand  left: * → %s", rs)
		}
		if _intersect(l[1:], r[j:], reason) {
			return true
		}
	}

	return false
}
