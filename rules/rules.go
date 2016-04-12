//rules package encloses the rules to be used in
//regexp.

package rules

const (
	EMAIL_RULE    = `^[a-zA-Z0-9.!#$%&\'*+\/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$`
	PIN_CODE_RULE = `^[1-9][0-9]{5}$`
	MOBILE_RULE   = `^[1-9]\d{9}$`
)
