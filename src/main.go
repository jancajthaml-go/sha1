package main

import "os"

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage : ./sha1 <input>\n"))
		return
	}

	sha1 := Digest([]byte(os.Args[1]))
	os.Stdout.Write(sha1)

	os.Exit(0)

	/*
		  Currently does not work
			   sha1 aaa -> 7e240de74fb1ed08fa08d38063f6a6a91462a815  SHOULD BE
			            -> 22d9c2107fd9877ef49fc3321ab02920bd5cf0c5  ACTUALLY IS
	*/
}
