package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type word struct {
	name   string
	amount int
}

type byAmount []word // to make custom sort custom slice is needed

func (a byAmount) Len() int { return len(a) }
func (a byAmount) Less(i, j int) bool { // by amount and name in 1 function
	if a[i].amount == a[j].amount {
		return a[i].name < a[j].name // sort names in order
	}
	return a[i].amount > a[j].amount // sort amount in reverse order
}
func (a byAmount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Top10(s string) []string {
	words := strings.Fields(s)
	countWords := make(map[string]int)
	for _, w := range words {
		w = strings.TrimSpace(w)
		if unicode.IsDigit(rune(w[0])) || unicode.IsSymbol(rune(w[0])) {
			continue
		}
		w = strings.ToLower(w)
		wRune := []rune(w)
		if len(wRune) > 0 && unicode.IsPunct(wRune[len(wRune)-1]) {
			wRune = wRune[:len(wRune)-1]
		}
		if len(wRune) > 0 && unicode.IsPunct(wRune[0]) {
			wRune = wRune[1:]
		}
		if len(wRune) == 0 {
			continue
		}
		w = string(wRune)
		countWords[w]++
	}
	res := make([]word, 0, len(countWords))
	for k, v := range countWords {
		res = append(res, word{name: k, amount: v})
	}
	sort.Sort(byAmount(res))
	result := make([]string, 0, 10)
	var l int
	if len(res) < 10 {
		l = len(res)
	} else {
		l = 10
	}
	for i := 0; i < l; i++ {
		result = append(result, res[i].name)
	}
	return result
}
