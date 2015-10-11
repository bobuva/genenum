package main

import (
	"flag"
	"fmt"
)

type Options struct {
	inFilename  string
	name        string
	numValues   int
	packageName string
	usePrefix   bool
}

// SetupOptions parses and post-normalizes options.
func SetupOptions() (options *Options) {
	options = new(Options)

	flag.StringVar(&options.inFilename, "inf", "", "Filename of file containing enumerated type string values")
	flag.StringVar(&options.name, "name", "", "Enumerated type name")
	flag.IntVar(&options.numValues, "numvalues", -1, "Number of enumerated type values")
	flag.StringVar(&options.packageName, "pkg", "main", "Package name. Defaults to 'main'")
	flag.BoolVar(&options.usePrefix, "useprefix", true, "If true, the name of the type will be used as a prefix with an underscore character")
	flag.Parse()

	if options.numValues > 200 {
		fmt.Printf("Warning: You have specified a value of %v for the number of values. Kinda large, don't you think?\n", options.numValues)
	}
	return
}
