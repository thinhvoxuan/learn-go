package main

import (
	"fmt"
	"sort"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type ArrayStringString [][]string

func (s ArrayStringString) Len() int {
	return len(s)
}
func (s ArrayStringString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ArrayStringString) Less(i, j int) bool {
	fArr := s[i]
	sArr := s[j]
	return fArr[len(fArr)-1] < sArr[len(sArr)-1]
}

func domainForwarding(redirects [][]string) (summaries [][]string) {
	mappingDomain := map[string][]string{}
	for _, redirect := range redirects {
		fr := redirect[0]
		to := redirect[1]
		found := false

		if mappingDomain[fr] != nil {
			mappingDomain[to] = mappingDomain[fr]
			delete(mappingDomain, fr)
		}

		for key, summary := range mappingDomain {
			if contains(summary, fr) {
				found = true
				mappingDomain[key] = append(summary, to)
				sort.Strings(mappingDomain[key])
			}
			if contains(summary, to) {
				found = true
				mappingDomain[key] = append(summary, fr)
				sort.Strings(mappingDomain[key])
			}
		}
		if !found {
			sort.Strings(redirect)
			mappingDomain[to] = redirect
		}
	}

	var keys []string
	for k := range mappingDomain {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		summaries = append(summaries, mappingDomain[k])
	}

	return
}

func domainType(domains []string) (domainList []string) {
	for _, domain := range domains {
		splitText := strings.Split(domain, ".")
		domainType := splitText[len(splitText)-1]
		domainTypeName := ""
		switch domainType {
		case "org":
			domainTypeName = "organization"
		case "com":
			domainTypeName = "commercial"
		case "net":
			domainTypeName = "network"
		case "info":
			domainTypeName = "information"
		}
		domainList = append(domainList, domainTypeName)
	}
	return
}

func woodBars(bars []int) int {
	sort.Ints(bars)
	if len(bars) == 1 {
		return bars[0]
	}

	if bars[0] == bars[1] {
		return woodBars(bars[1:])
	}

	return woodBars(reduceValue(bars))
}

func reduceValue(bars []int) (result []int) {
	fValue := bars[0]
	result = []int{fValue}
	for _, v := range bars[1:] {
		if v%fValue != 0 {
			result = append(result, v%fValue)
		}
	}
	return
}

func main() {
	text := "en.wiki.org"

	fmt.Printf("domain: ", domain)
}
