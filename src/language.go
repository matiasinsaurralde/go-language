package language

import ( "fmt"
	 "strings" )

func LoadText( fileName string ) string {

	fmt.Printf( "loading file '%s'.\n", fileName )

	return "this is a sample sentence"
}

func ExtractLetters( text string ) {

	splits := strings.Split( text, "" )

	fmt.Printf( "original string: '%s'\n", text )

	fmt.Printf( "splits: %q\n", splits )
}
