package language

import(	"strings"
	"io/ioutil"
	"sort"
	"fmt" )

func ReadText( fileName string ) string {
	fileContent, _ := ioutil.ReadFile( fileName )
	return string( fileContent )
}

func ExtractLetters( text string ) map[string]int {

	letters := make( map[string]int )

	splits := strings.Split( strings.ToLower( text ), "" )

	for i := 0; i < len( splits ); i++ {
		letters[ splits[ i ] ] += 1
	}

	/* 

		var k string
		var v interface {}

		for k, v = range letters {
			fmt.Printf( "%s => %d\n", k, v )
		}

	*/

	return letters

}

func BuildModel( corpusFilename string, modelName string  ) {

	corpus := ReadText( corpusFilename )

	output := ExtractLetters( corpus )
	ordered := OrderPlease( output )

	var k string
	var v interface {}

	for k, v = range ordered {
		 fmt.Printf( "%s => %d\n", k, v )
	}

	// fmt.Printf("%s\n", corpus)

}

func OrderPlease( letters map[string]int ) map[string]int {

	var keys []int
	sw := make( map[int]string )
	fm := make ( map[string]int )

	for k, v := range letters {
		sw[ v ] = k
	}

	for k := range sw {
		keys = append( keys, k )
	}

	sort.Ints( keys )

	for _, k := range keys {
		// fmt.Printf("key: %d, Value: %s", k, sw[ k ] )
		fm[ sw[k] ] = 0
	}

	for k,v := range fm {
		fmt.Printf("%s => %d\n", k, v )
	}

	fmt.Printf("***\n")

	return letters
}
