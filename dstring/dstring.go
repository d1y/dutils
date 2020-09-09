// Author: d1y<chenhonzhou@gmail.com>

package dstring

// TODO
// const (
// 	Symbol1  = "!"
// 	Symbol2  = "@"
// 	Symbol3  = "#"
// 	Symbol4  = "$"
// 	Symbol5  = "%"
// 	Symbol6  = "^"
// 	Symbol7  = "&"
// 	Symbol8  = "*"
// 	Symbol9  = "("
// 	Symbol10 = ")"
// 	Symbol11 = "~"
// 	Symbol12 = "`"
// )

// Padend padding end(right)
//
// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd
func Padend(targetString string, paddingString string, length int) string {
	var appendString = Fill(paddingString, length)
	var result = targetString + appendString
	return result
}

// Padstart padding start(left)
//
// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padStart
func Padstart(targetString string, paddingString string, length int) string {
	var appendString = Fill(paddingString, length)
	var result = appendString + targetString
	return result
}

// Padding padding string(left and right)
func Padding(targetString string, paddingString string, length int) string {
	var appendString = Fill(paddingString, length)
	var result = appendString + targetString + appendString
	return result
}

// Fill fill string
func Fill(paddingString string, length int) string {
	var r string
	for i := 0; i < length; i++ {
		r += paddingString
	}
	return r
}
