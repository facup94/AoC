package day7

import (
	"strings"
)

type ip struct {
	supernetSequences []string
	hypernetSequences []string
}

func (i ip) hasABBAOnSupernet() bool {
	for _, sequence := range i.supernetSequences {
		if hasABBA(sequence) {
			return true
		}
	}
	return false
}

func (i ip) hasABBAOnHypernet() bool {
	for _, sequence := range i.hypernetSequences {
		if hasABBA(sequence) {
			return true
		}
	}
	return false
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if isABBA(s[i : i+4]) {
			return true
		}
	}
	return false
}

func isABBA(s string) bool {
	return s[0] != s[1] && s[0] == s[3] && s[1] == s[2]
}

func (i ip) getABAs() []string {
	var abas []string
	for _, s := range i.supernetSequences {
		for x := 0; x < len(s)-2; x++ {
			if isABA(s[x : x+3]) {
				abas = append(abas, s[x:x+3])
			}
		}
	}
	return abas
}

func (i ip) getBABs() []string {
	var babs []string
	for _, s := range i.hypernetSequences {
		for x := 0; x < len(s)-2; x++ {
			if isABA(s[x : x+3]) {
				babs = append(babs, s[x:x+3])
			}
		}
	}
	return babs
}

func isABA(s string) bool {
	return s[0] != s[1] && s[0] == s[2]
}

func stringToIP(s string) ip {
	ss := strings.FieldsFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})

	var supernetSequences, hypernetSequences []string
	for i, s2 := range ss {
		if i%2 == 0 {
			supernetSequences = append(supernetSequences, s2)
		} else {
			hypernetSequences = append(hypernetSequences, s2)
		}
	}

	return ip{supernetSequences: supernetSequences, hypernetSequences: hypernetSequences}
}
