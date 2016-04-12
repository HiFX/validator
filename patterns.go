//validator package implements commonly used validation functions

package validator

import (
	"github.com/hifx/validator/rules"
	"regexp"
)

func IsEmailValid(email string) bool {
	emailRegexp, _ := regexp.Compile(rules.EMAIL_RULE)
	return emailRegexp.MatchString(email)
}

func IsMobileNumberValid(mobile string) bool {
	mobileRegExp, _ := regexp.Compile(rules.MOBILE_RULE)
	return mobileRegExp.MatchString(mobile)
}

func IsPinCodeValid(pinCode string) bool {
	pinRegexp, _ := regexp.Compile(rules.PIN_CODE_RULE)
	return pinRegexp.MatchString(pinCode)
}
