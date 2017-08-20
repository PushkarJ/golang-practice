ckage main
import (
	"fmt"
	"strings"
)

func main() {
	input := "woahwoahwoahwoahwoahwoahwoowoahwoah"
	fmt.Println(strings.Count(input, "woah"))
	fmt.Println(strings.Count(input, "woo"))
	fmt.Println(strings.ContainsAny(input, "eiu"))
	fmt.Println(strings.ContainsAny(input, "ao"))
	fmt.Println(strings.Compare(input, "woah"))
	fmt.Println(strings.EqualFold(input, "Woahwoahwoahwoahwoahwoahwoowoahwoah"))
	fmt.Println(strings.HasPrefix(input, "woah"))
	fmt.Println(strings.HasSuffix(input, "woah"))
	fmt.Println(strings.Index(input, "woo"))
	fmt.Println(strings.IndexAny(input, "eiu"))
	fmt.Println(strings.IndexAny(input, "ao"))

}

