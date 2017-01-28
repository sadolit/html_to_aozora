package main

import "regexp"

func RubyToBunko(in string) string {
	var removeRp = regexp.MustCompile("(?U)<rp>(.*)<\\/rp>")
	var removeRubyAndRbTags = regexp.MustCompile("(<ruby>|<\\/ruby>|<rb>|<\\/rb>)")
	var removeRtStart = regexp.MustCompile("<rt>")
	var removeRtEnd = regexp.MustCompile("</rt>")
	var woRp = removeRp.ReplaceAllString(in, "")
	var woRubtAndRb = removeRubyAndRbTags.ReplaceAllString(woRp, "")
	var woRtStart = removeRtStart.ReplaceAllString(woRubtAndRb, "《")
	return removeRtEnd.ReplaceAllString(woRtStart, "》")
}
