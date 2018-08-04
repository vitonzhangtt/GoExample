package main

import "fmt"
import "net"
import "net/url"

/*
1. Link
https://gobyexample.com/url-parsing
*/

func main() {

/*
2. We’ll parse this example URL, which includes a scheme, 
	authentication info, host, port, path, query params, 
	and query fragment.
*/
	s := "postgres://user:pass@host.com:5432/path?k=value;k2=value2#f"
	fmt.Println(s)	

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
/*
3. User contains all authentication info; call Username and 
	Password on this for individual values.
*/
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	
	fmt.Println("---------------------")	
	p, _ := u.User.Password()
	fmt.Println(p)
	
/*
4. The Host contains both the hostname and the port, if present. 
	Use SplitHostPort to extract them.
*/
	fmt.Println("---------------------")
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	
/*
5. Here we extract the path and the fragment after the #.
*/
	fmt.Println("---------------------")
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	
/*
6. To get query params in a string of k=v format, use RawQuery. 
	You can also parse query params into a map. The parsed 
	query param maps are from strings to slices of strings, 
	so index into [0] if you only want the first value.
*/
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])	
}





