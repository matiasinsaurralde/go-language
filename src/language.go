package language

import( "encoding/json"
	"path/filepath"
	"io/ioutil"
	"strings"
	"sort"
	"fmt"
	"os" )

type Letter struct {

	Name	string
	Freq	float64	`json:",string"`

}

func ReadText( fileName string ) string {
	fileContent, _ := ioutil.ReadFile( fileName )
	return string( fileContent )
}

func ExtractLetters( text string ) map[string]float64 {

	abc := []string{ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" }

	letters := make( map[string]float64 )

	totalLetters := 0.0

	var k string
	var v float64

	splits := strings.Split( strings.ToLower( text ), "" )

	for i := 0; i < len( splits ); i++ {

		index := sort.SearchStrings( abc, splits[i] )

		if ( splits[i] == abc[index] ) {	// by now we'll just use abc.

			letters[ splits[ i ] ] += 1
			totalLetters++

		}
	}

	for k, v = range letters {
		letters[ k ] = v/totalLetters*100.0
	}

	return letters

}

func BuildModel( corpusFilename string, modelName string  ) {

	corpus := ReadText( corpusFilename )

	output := ExtractLetters( corpus )
	// ordered := OrderPlease( output )

	/*var k string
	var v interface {}

	for k, v = range output {
		 fmt.Printf( "%s => %f\n", k, v )
	}*/

	var jsonData []byte
	jsonData, _ = json.Marshal( output )

	ioutil.WriteFile( "models/"+modelName, jsonData, 0644 )

}

func LoadModels() {
	fmt.Printf("loading models.\n")

	filepath.Walk("./models/", func(path string, _ os.FileInfo, _ error) {
		//fmt.Println(path)
	})
	
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

/*	for k,v := range fm {
		fmt.Printf("%s => %d\n", k, v )
	}*/

	fmt.Printf("***\n")

	return letters
}
