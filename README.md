# genenum
Generate Go code for a new enumerated type.

## Purpose
Go does not support enumerated types but makes it very easy to simulate them using constants and the iota enumerator. Each time one wants to have a so-called enumerated type, however, it is necessary to create a new set of constants along with string mapping.

I decided to make this easier by writing this little utility for generating the Go code based on the little bit of information about the type you want to create.

## Usage
There are two input modes:
    1. using an input file definition of the enumerated type strings, and
    2. using just the number of values for the enumerated type.

### Input File method
Create a new text file whose name will be used as the enumerated type name.
Each line of this file should be a separate string name for an enumerated type. The Genre file is an example of an input file.

To generate the Go code from an input file pass the name of file to the program using the -inf command-line option as shown below.

    ./genenum -inf=Genre

This will generate a Genre.go file.

### Other Method for Generating File
The other option support by genenum is to pass the name of the enumerated type and the number of values to generate. Here is an example:

    ./genenum -name=Colors -numvalues=8

This will generate a Colors.go file that uses the name of the type and an integer to create temporary enumerated type names which you can change.


### Common Command-Line Options
genenum supports other options for either of the above methods:

##### Prefix
genenum takes an optional -prefix parameter that specifies whether to prefix the enumerated type values with the name of the type and an underscore. Instead of Red, Green, Blue in an enumerated type Color, for example, the values would appear as Color_Red, Color_Green, Color_Blue. The string representation of each enumerated value, however, would not include the prefix. The purpose of this option is to make it easy to avoid name conflicts.

The default value for this option is true.

Example:
    ./genenum -inf=Genre -useprefix=false 
    
##### Package Name
genenum takes an optional -pkg parameter for specifying the package name to be put in the generated file. If this option is not provided, genenum will use 'main'.

Example:
    ./genenum -inf=Genre -pkg=music

