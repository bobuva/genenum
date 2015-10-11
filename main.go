package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
	indent = "  "
)

type EnumData struct {
	typeName      string
	useValueNames bool
	valueNames    []string
	numberValues  int
	packageName   string
}

func main() {
	options := SetupOptions()

	if options.inFilename == "" { // then requires name and numValues
		if options.name == "" || options.numValues < 1 {
			usage()
			return
		}
	}

	filename, err := generateEnum(options)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("Generated enumeration in %v\n", filename)
}

func usage() {
	fmt.Printf("Usage: %v [-inf=<input filename>] | [[-name=<enumerated type name>] [-numvalues=<number of values>]] [-pkg=<package name - defaults to 'main'>] [-prefix=<boolean>]\n", os.Args[0])
	fmt.Printf("ex: %v -inf=Colors \n", os.Args[0])
	fmt.Printf("ex: %v -name=Colors -numvalues=4\n", os.Args[0])
	fmt.Printf("    %v -name=Languages -numvalues=8 -pkg=words\n")
	fmt.Printf("ex: %v -inf=Colors -pkg=paint\n", os.Args[0])
	fmt.Printf("ex: %v -inf=Colors -prefix=true")
}

// creates a new enumerated type in a file named <name>.go
// and returns the full path to the file.
func generateEnum(options *Options) (string, error) {
	enumeration := new(EnumData)

	// if there's an input file, then the file's base name is the name of
	// the enumerated type, and each line of the file contains a string
	// representation of an enumerated type value.
	//
	if options.inFilename != "" {
		file, err := os.Open(options.inFilename)
		if err != nil {
			return "", err
		}
		defer file.Close()

		numberLines := 0
		typeValues := make([]string, 0)
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			typeValues = append(typeValues, string(line))
			numberLines++
		}

		enumeration.typeName = filepath.Base(options.inFilename)
		enumeration.numberValues = numberLines
		enumeration.valueNames = typeValues
		enumeration.packageName = options.packageName
		enumeration.useValueNames = true
		filename, err := internalGenerateEnum(enumeration, options)
		if err != nil {
			return "", err
		} else {
			return filename, nil
		}
	} else {
		// Generating generic enumerated type
		enumeration.typeName = options.name
		enumeration.numberValues = options.numValues
		enumeration.packageName = options.packageName
		enumeration.useValueNames = false
		filename, err := internalGenerateEnum(enumeration, options)

		if err != nil {
			return "", err
		} else {
			return filename, nil
		}
	}

	return "", nil
}

func internalGenerateEnum(enum *EnumData, options *Options) (string, error) {
	lines := make([]string, 0, (enum.numberValues*2)+20)
	lines = append(lines, fmt.Sprintf("package %v\n\n", options.packageName))

	// prefixing of the type name and an underscore
	prefix := ""
	if options.usePrefix {
		prefix = enum.typeName + "_"
	}
	// enumeration value type: int8 or int16
	var intSize string
	if enum.numberValues < 256 {
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
	lines = append(lines, fmt.Sprintf("type %v %v ", enum.typeName, intSize))
	lines = append(lines, "const (")
	lines = append(lines, fmt.Sprintf("%v_%v = iota", indent, enum.typeName))

	for i := 1; i <= enum.numberValues; i++ {
		if enum.useValueNames {
			lines = append(lines, fmt.Sprintf("%v%v%v", indent, prefix, enum.valueNames[i-1]))
		} else {
			lines = append(lines, fmt.Sprintf("%v%v%v%v", indent, prefix, enum.typeName, i))
		}
	}
	lines = append(lines, ")")

	// var nameStrings =
	// map[string]Name {
	//   "Name1":	Name1,
	//	 ...
	//	 "NameN":	NameN,
	// }
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("var %vStrings = ", enum.typeName))
	lines = append(lines, fmt.Sprintf("map[string]%v {", enum.typeName))
	for i := 1; i <= enum.numberValues; i++ {
		if enum.useValueNames {
			lines = append(lines, fmt.Sprintf("%v\"%v\":\t%v%v,", indent, enum.valueNames[i-1], prefix, enum.valueNames[i-1]))
		} else {
			lines = append(lines, fmt.Sprintf("%v\"%v%v\":\t%v%v%v,", indent, enum.typeName, i, prefix, enum.typeName, i))
		}
	}
	lines = append(lines, "}")

	// func (l Cost) String() string {
	//  for s, v := range costStrings {
	//    if l == v {
	//       return s
	//    }
	//  }
	//  return "invalid"
	// }
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("func (l %v) String() string {", enum.typeName))
	lines = append(lines, fmt.Sprintf("%vfor s, v := range %vStrings {", indent, enum.typeName))
	lines = append(lines, fmt.Sprintf("%v%vif l == v {", indent, indent))
	lines = append(lines, fmt.Sprintf("%v%v%vreturn s", indent, indent, indent))
	lines = append(lines, fmt.Sprintf("%v%v}", indent, indent))
	lines = append(lines, fmt.Sprintf("%v}", indent))
	lines = append(lines, fmt.Sprintf("%vreturn \"invalid\"", indent))
	lines = append(lines, fmt.Sprintf("}"))

	file, err := os.Create(enum.typeName + ".go")
	if err != nil {
		return "", err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	w.Flush()
	return file.Name(), nil
}
