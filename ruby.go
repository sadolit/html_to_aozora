package main

import "regexp"

//RubyToBunko Replaces HTML ruby tags (<ruby>, <rb>, <rt>) with Aozora Bunko tags 《》
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

//RemoveHTMLTags Removes all HTML tags
func RemoveHTMLTags(in string) string {
	var removeTag = regexp.MustCompile("(?Ui)<(.|\n)*>")
	return removeTag.ReplaceAllString(in, "")
}

//AddNewline Replaces <br> tag with newline
func AddNewline(in string) string {
	var convertBrToN = regexp.MustCompile("<\\/{0,1}\\s*br\\s*/{0,1}>")
	return convertBrToN.ReplaceAllString(in, "\n")
}
