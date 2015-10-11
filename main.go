package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	indent = "  "
)

func main() {
	options := SetupOptions()

	if options.name == "" || options.numValues < 1 {
		usage()
		return
	}

	filename, err := generateEnum(options)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("Generated enumeration in %v\n", filename)
}

func usage() {
	fmt.Printf("Usage: %v -name=<enumerated type name> -numvalues=<number of values> -pkg=<package name - defaults to 'main'>\n", os.Args[0])
	fmt.Printf("ex: %v -name=Colors -numvalues=4\n", os.Args[0])
	fmt.Printf("    %v -name=Languages -numvalues=8 pkg=words\n")
}

// creates a new enumerated type in a file named <name>.go
// and returns the full path to the file.
func generateEnum(options *Options) (string, error) {
	file, err := os.Create(options.name + ".go")
	if err != nil {
		return "", err
	}
	defer file.Close()

	lines := make([]string, 0, 100)
	lines = append(lines, fmt.Sprintf("package %v\n\n", options.packageName))

	// enumeration value type: int8 or int16
	var intSize string
	if options.numValues < 256 {
		intSize = "int8"
	} else {
		intSize = "int16"
	}

	// type name intSize
	// const (
	//  _name = iota
	//  name1
	//	...
	//  nameN
	// )
	//
	lines = append(lines, fmt.Sprintf("type %v %v ", options.name, intSize))
	lines = append(lines, "const (")
	lines = append(lines, fmt.Sprintf("%v_%v = iota", indent, options.name))
	for i := 1; i <= options.numValues; i++ {
		lines = append(lines, fmt.Sprintf("%v%v%v", indent, options.name, i))
	}
	lines = append(lines, ")")

	// var nameStrings =
	// map[string]Name {
	//   "Name1":	Name1,
	//	 ...
	//	 "NameN":	NameN,
	// }
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("var %vStrings = ", options.name))
	lines = append(lines, fmt.Sprintf("map[string]%v {", options.name))
	for i := 1; i <= options.numValues; i++ {
		lines = append(lines, fmt.Sprintf("%v\"%v%v\":\t%v%v,", indent, options.name, i, options.name, i))
	}
	lines = append(lines, "}")

	//func (l Cost) String() string {
	//  for s, v := range costStrings {
	//    if l == v {
	//       return s
	//    }
	//  }
	//  return "invalid"
	//}
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("func (l %v) String() string {", options.name))
	lines = append(lines, fmt.Sprintf("%vfor s, v := range %vStrings {", indent, options.name))
	lines = append(lines, fmt.Sprintf("%v%vif l == v {", indent, indent))
	lines = append(lines, fmt.Sprintf("%v%v%vreturn s", indent, indent, indent))
	lines = append(lines, fmt.Sprintf("%v%v}", indent, indent))
	lines = append(lines, fmt.Sprintf("%v}", indent))
	lines = append(lines, fmt.Sprintf("%vreturn \"invalid\"", indent))
	lines = append(lines, fmt.Sprintf("}"))

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	w.Flush()
	return file.Name(), nil
}

// func writeLines(lines []string, path string) error {

// 	w := bufio.NewWriter(file)
// 	for _, line := range lines {
// 		fmt.Fprintln(w, line)
// 	}
// 	return w.Flush()
// }
