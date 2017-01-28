package main

import (
	"fmt"
	"testing"
	"time"
)

var testRubyToBunkoData = []struct {
	in       string
	expected string
}{
	{"", ""},
	{"<ruby><rb>媒妁人</rb><rp>（</rp><rt>なかうど</rt><rp>）</rp></ruby>",
		"媒妁人《なかうど》"},
	{"<ruby><rb>媒妁人</rb><rp>（</rp><rt>なかうど</rt><rp>）</rp></ruby><ruby><rb>先</rb><rp>（</rp><rt>ま</rt><rp>）</rp></ruby>",
		"媒妁人《なかうど》先《ま》"},
	{"<ruby><rb>媒妁人</rb><rp>（</rp><rt>なかうど</rt><rp>）</rp></ruby><ruby><rb>先</rb><rp>（</rp><rt>ま</rt><rp>）</rp></ruby>づいふめでたしと、<ruby><rb>舅姑</rb><rp>（</rp><rt>きうこ</rt><rp>）</rp></ruby>またいふめでたしと、親類等皆いふめでたしと、<ruby><rb>知己</rb><rp>（</rp><rt>ちき</rt><rp>）</rp></ruby><ruby><rb>朋友</rb><rp>（</rp><rt>ほういう</rt><rp>）</rp></ruby>皆いふめでたしと、<ruby><rb>渠等</rb><rp>（</rp><rt>かれら</rt><rp>）</rp></ruby>は<ruby><rb>欣々然</rb><rp>（</rp><rt>きん／＼ぜん</rt><rp>）</rp></ruby>として新夫婦の婚姻を祝す、婚礼果してめでたきか。<br>",
		"媒妁人《なかうど》先《ま》づいふめでたしと、舅姑《きうこ》またいふめでたしと、親類等皆いふめでたしと、知己《ちき》朋友《ほういう》皆いふめでたしと、渠等《かれら》は欣々然《きん／＼ぜん》として新夫婦の婚姻を祝す、婚礼果してめでたきか。<br>"},
	{"ABC<ruby><rb>媒妁人</rb><rp>（</rp><rt>なかうど</rt><rp>）</rp></ruby>ABC",
		"ABC媒妁人《なかうど》ABC"},
	{"<ruby>攻殻<rt>こうかく</rt>機動隊<rt>きどうたい</rt></ruby>",
		"攻殻《こうかく》機動隊《きどうたい》"},
}

var TestRemoveHtmlTagsData = []struct {
	in       string
	expected string
}{
	{"", ""},
	{"<div>", ""},
	{"</div>", ""},
	{"<div></div>", ""},
	{`<div class="tab0">CSS code formatter</div><div class="tab2">CSS code  compressor</div><br></br>`, "CSS code formatterCSS code  compressor"},
	{`合<rb>合<img gaiji="gaiji" src="../../../gaiji/2-03/2-03-54.png" alt="※(「丞／（厄－厂）」、第4水準2-3-54)" class="gaiji"></rb><rp>`,
		"合合"},
}

var TestAddNewlineData = []struct {
	in       string
	expected string
}{
	{"", ""},
	{"<div>", "<div>"},
	{"abc<br>abc", "abc\nabc"},
	{"abc</br>abc", "abc\nabc"},
	{"abc<br/>abc", "abc\nabc"},
	{"abc</  br>abc", "abc\nabc"},
	{"abc< br />abc", "abc\nabc"},
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func TestRubyToBunko(t *testing.T) {
	for _, testData := range testRubyToBunkoData {
		beginTime := makeTimestamp()
		actualBunko := RubyToBunko(testData.in)
		if testData.expected != actualBunko {
			t.Errorf("Input: [%s].Expected [%s], but it was [%s] instead.", testData.in, testData.expected, actualBunko)
		}
		endTime := makeTimestamp()
		fmt.Printf("Test took %d\n", (endTime - beginTime))
	}
}

func TestRemoveHtmlTags(t *testing.T) {
	for _, testData := range TestRemoveHtmlTagsData {
		beginTime := makeTimestamp()
		actualBunko := RemoveHtmlTags(testData.in)
		if testData.expected != actualBunko {
			t.Errorf("Input: [%s].Expected [%s], but it was [%s] instead.", testData.in, testData.expected, actualBunko)
		}
		endTime := makeTimestamp()
		fmt.Printf("Test took %d\n", (endTime - beginTime))
	}
}

func TestAddNewline(t *testing.T) {
	for _, testData := range TestAddNewlineData {
		beginTime := makeTimestamp()
		actualBunko := AddNewline(testData.in)
		if testData.expected != actualBunko {
			t.Errorf("Input: [%s].Expected [%s], but it was [%s] instead.", testData.in, testData.expected, actualBunko)
		}
		endTime := makeTimestamp()
		fmt.Printf("Test took %d\n", (endTime - beginTime))
	}
}
