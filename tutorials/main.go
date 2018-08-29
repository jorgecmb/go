package main

import (
	"fmt"

	"github.com/jorgecmb/tutorials/basictypes"
	"github.com/jorgecmb/tutorials/constants"
	"github.com/jorgecmb/tutorials/constantsnumeric"
	"github.com/jorgecmb/tutorials/conversions"
	"github.com/jorgecmb/tutorials/exportednames"
	"github.com/jorgecmb/tutorials/functions"
	"github.com/jorgecmb/tutorials/functionscont"
	"github.com/jorgecmb/tutorials/imports"
	"github.com/jorgecmb/tutorials/multipleresults"
	"github.com/jorgecmb/tutorials/namedresults"
	"github.com/jorgecmb/tutorials/packages"
	"github.com/jorgecmb/tutorials/typeinference"
	"github.com/jorgecmb/tutorials/variables"
	"github.com/jorgecmb/tutorials/variablesinitializers"
	"github.com/jorgecmb/tutorials/variablesinitializersshort"
	"github.com/jorgecmb/tutorials/zero"
)

func main() {
	fmt.Println("**** STARTING MAIN ****")

	fmt.Printf("\npackages: ")
	packages.Run()

	fmt.Printf("imports: ")
	imports.Run()

	fmt.Printf("exported-names: ")
	exportednames.Run()

	fmt.Printf("functions: ")
	functions.Run()

	fmt.Printf("functions-cont: ")
	functionscont.Run()

	fmt.Printf("multiple-results: ")
	multipleresults.Run()

	fmt.Printf("named-results: ")
	namedresults.Run()

	fmt.Printf("variables: ")
	variables.Run()

	fmt.Printf("variables-initializers: ")
	variablesinitializers.Run()

	fmt.Printf("variables-initializers-short: ")
	variablesinitializersshort.Run()

	fmt.Printf("basic-types: ")
	basictypes.Run()

	fmt.Printf("zero: ")
	zero.Run()

	fmt.Printf("conversions: ")
	conversions.Run()

	fmt.Printf("typeinference: ")
	typeinference.Run()

	fmt.Printf("constants: ")
	constants.Run()

	fmt.Printf("constants-numeric: ")
	constantsnumeric.Run()

	fmt.Println("\n**** FINISING MAIN ****")
}
