package appstore

import "errors"

var _EVN = map[string]int8{
	"production": PRODUCTION,
	"sandbox":    SANDBOX,
}
var RESP_NIL = errors.New("Response is nil .")
var GP_PUBLIC_KEY = "xxxxxxxxxxxx"

const (
	_ = int8(iota)
	PRODUCTION
	SANDBOX
)

// APPLE PAY URL
const (
	APPLE_VERIY_PRODUCTION_URL = "https://sandbox.itunes.apple.com/verifyReceipt"
	APPLE_VERIY_SANDBOX_URL    = "https://sandbox.itunes.apple.com/verifyReceipt"
)
