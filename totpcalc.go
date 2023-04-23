package main

import (
	"flag"
	auth "github.com/dgryski/dgoogauth"
	"time"
)

var Time int64
var Secret string
var Int int64

func GenerateTOTP(secret string, time int64) int {
	return auth.ComputeCode(secret, time)
}

func init() {
	flag.Int64Var(&Time, "time", time.Now().UTC().Unix(), "Time to use, in seconds since the epoch unix time")
	flag.Int64Var(&Int, "interval", 30, "Number of digits to use on the token")
	flag.StringVar(&Secret, "secret", "", "Secret to use")
	flag.Parse()
}

func main() {
	println("Digit: ", GenerateTOTP(Secret, Time/Int))
}
