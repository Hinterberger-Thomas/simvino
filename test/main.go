package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {

	checksum := sha512.Sum512([]byte("klajsdklöfjkalsdfjklöadsjklfkjls<dflökds<jköfkdlöfjklösdaöjklfjklodsafjkweqipofjqdweiofjsadocn3eihjfnrdefvhdsoifcmnjdwiochnq3ebnfcowdbnfvhnreqfjewoidjdwqofcmnsadufgijh93efjdcdsoifhreojf"))
	sum := string(checksum[:])
	fmt.Println(string(sum))
	sum2 := string(sha256.Sum256([]byte("jskdfjadslfkajskfjasdölfajsdflköasjklf"))[:])
	fmt.Println(sum2)
}
