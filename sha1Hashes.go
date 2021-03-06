package main

import "crypto/sha1"
import "fmt"

/*
1. Link
https://gobyexample.com/sha1-hashes
*/

/*
2.
SHA1 hashes are frequently used to compute short identities for binary 
	or text blobs. For example, the git revision control system 
	uses SHA1s extensively to identify versioned files and 
	directories.
*/

/*
3. Go implements several hash functions in various crypto/* packages.
*/
func main() {
	s := "sha1 this string"
/*
4. The pattern for generating a hash is sha1.New(), sha1.Write(bytes), 
	then sha1.Sum([]byte{}). Here we start with a new hash.
*/
	h := sha1.New()
/*
5. Write expects bytes. If you have a string s, use []byte(s) to 
	coerce it to bytes.
*/
	h.Write([]byte(s))
	
/*
6. This gets the finalized hash result as a byte slice. The 
	argument to Sum can be used to append to an existing 
	byte slice: it usually isn’t needed.
*/
	bs := h.Sum(nil)
	
	fmt.Println(s)
/*
7. Use the %x format verb to convert a hash results to a hex string.
*/
	fmt.Printf("%x\n", bs)
}

/*
8. For example, to compute MD5 hashes import crypto/md5 and use md5.New().
*/


