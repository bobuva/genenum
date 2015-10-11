package main

import (
	"flag"
	"fmt"
)

type Options struct {
	name        string
	numValues   int
	packageName string
}

// SetupOptions parses and post-normalizes options.
func SetupOptions() (options *Options) {
	options = new(Options)

	flag.StringVar(&options.name, "name", "", "Enumerated type name")
	flag.IntVar(&options.numValues, "numvalues", -1, "Number of enumerated type values")
	flag.StringVar(&options.packageName, "pkg", "main", "Package name. Defaults to 'main'")
	flag.Parse()

	if options.numValues > 200 {
		fmt.Printf("Warning: You have specified a value of %v for the number of values. Kinda large, don't you think?\n", options.numValues)
	}
	return
}
