package validator

import (
	"fmt"
	"regexp"
)

// 表达式文档 https://c.runoob.com/front-end/854/

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}

// ValidatorPhoneNumber 匹配纯粹的11位数字
func ValidatorPhoneNumber(phone string) bool {
	matched, err := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", phone)
	if err != nil {
		fmt.Print(err)
		return false
	}
	return matched
}
