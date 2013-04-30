package main

import(	"language"
	"fmt"
	"os" )

func main() {

	fmt.Printf( "building model '%s'\n", os.Args[2] )

	language.BuildModel( os.Args[1], os.Args[2] )

}
