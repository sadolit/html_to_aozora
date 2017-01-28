package main

import "regexp"

func RubyToBunko(in string) string {
	var getRubyAndRt = regexp.MustCompile("(?Ui)<ruby>.*<rb>(.*)<\\/rb>.*<rt>(.*)<\\/rt>.*<\\/ruby>")
	var results = getRubyAndRt.ReplaceAllString(in, "$1《$2》")
	return results
}
