package main

import (
"fmt"
"strings"
)

func main() {
test := []string{"I", "am", "stupid", "and", "weak"} // The slice of data
/* fmt.Println(test)
fmt.Println(len(test)) */
for i := 0; i < len(test); i++ {
// fmt.Println(i)
if test[i] == "stupid" {
test[i] = "smart"
} else if test[i] == "weak" {
test[i] = "strong"
}
}
semiformat := fmt.Sprintf("%q\n", test)  // Turn the slice into a string that looks like ["one" "two" "three"]
tokens := strings.Split(semiformat, " ") // Split this string by spaces
fmt.Printf(strings.Join(tokens, ", "))   // Join the Slice together (that was split by spaces) with commas
}
