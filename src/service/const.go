package main

const (
	SERVICE = "IAP SERVICE"
)

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
