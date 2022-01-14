package lib

import (
	"flag"
)

var (
	TestFlag bool
	GroupID  string
	Cve      bool
)

func init() {
	flag.BoolVar(&TestFlag, "t", false, "enable testing")
	flag.StringVar(&GroupID, "g", "", "qq group for testing")
	flag.BoolVar(&Cve, "nc", false, "no cve list")
}
