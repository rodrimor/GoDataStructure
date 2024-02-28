// Hands on exercise display emojis and literal strings

package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("Hello world %v\n", emoji.GlobeShowingAmericas)
	fmt.Println(`
In the hiding
hour of autophagy
ghosts hang out all day and talk to us.
An archival haunting demanding tribute:
     half a lime for breakfast every day.
بشرٌ يئنّونَ من الألمِ
human voices keening in pain
تُشعلُ أجسادَهُمَ النارُ
their bodies, consumed by fire
light up the...
`)
}
