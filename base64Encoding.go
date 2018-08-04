package main

/*
2. This syntax imports the encoding/base64 package with the 
	b64 name instead of the default base64. It’ll save us 
	some space below.
*/
import b64 "encoding/base64"
import f "fmt"

/*
1. Link
https://gobyexample.com/base64-encoding
*/

func main() {

/*
3. Here’s the string we’ll encode/decode.
*/
	data := "abc123!?$*&()'-=@~"
	
/*
4. The encoder requires a []byte so we cast our string to that type.
*/
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	f.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	f.Println(string(sDec))
	f.Println()
	
	f.Println("------------------------")
/*
5. This encodes/decodes using a URL-compatible base64 format.
*/
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	f.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	f.Println(string(uDec))

}
