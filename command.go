package main

import (
	"flag"
	"io/ioutil"
)

func defaultFlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)

	// Set the default Usage to empty
	f.Usage = func() {}

	return f
}
