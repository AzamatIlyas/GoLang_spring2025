package practice1_unpack_string

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcStr string) (string, error) {

	var flag bool

	if srcStr == "" {
		return "", nil
	}

	var res string
	var prevRune rune
	runes := []rune(srcStr)

	// for i, r := range runes {

	// 	if unicode.IsDigit(r){
	// 		if i == 0 && !flag{
	// 			return "",ErrInvalidString 
	// 		}

	// 		repeatCount := int(r - '0')

	// 		if i+1 < len(runes) {
	// 			if unicode.IsDigit(r) {
	// 				return "", ErrInvalidString	
	// 			}
	// 		}

	// 		if repeatCount == 0{
	// 			if len(runes) > 0{
	// 				res = res[:len(res)-1]
	// 			}
	// 		}else{
	// 			for i:=1; i < repeatCount; i++{
	// 				res = append(res, prevRune)
	// 			}
	// 		}
	// 	}else{
	// 		res = append(res, r)
	// 		prevRune = r
	// 		flag = true
	// 	}

	// }
	

	for i, r := range srcStr {
		if unicode.IsDigit(r) {
			if i == 0 || !flag {
				return srcStr, ErrInvalidString
			}

			repeatCount := int(r - '0')

			if i+1 < len(srcStr) && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}

			if repeatCount == 0 {
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
			} else {
				res += strings.Repeat(string(prevRune), repeatCount-1)
			}
		} else {
			res += string(r)
			prevRune = r
			flag = true
		}
	}

	return string(res), nil
}
