Program1

package main
import (
	"fmt"
)
func main() {
	var n bool
	fmt.Printf("%v,%T\n", n, n)
}
output:
false,bool



Program2

package main
import (
	"fmt"
)
func main() {
	var n uint16 = 42
	fmt.Printf("%v,%T\n", n, n)
}
output:
42,uint16


program3

package main
import (
	"fmt"
)
func main() {
	a:= 10
	b:= 8
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a%b)
	fmt.Println(a/b)
	}
output:
18
2
80
2
1	


	
program4	

	package main
import (
	"fmt"
)
func main() {
	var a int = 10
	var b int8 = 3
	fmt.Println(a+int(b))

	}
output:
13	
	
	
	
Program5
	
package main
import (
	"fmt"
)

func main() {
	a:= 10
	b:= 8
	fmt.Println(a&b)
	fmt.Println(a|b)
	fmt.Println(a^b)
	fmt.Println(a&^b)
	}
	
output:
8
10
2
2
	
	
	
	
Program6
		
package main
import (
	"fmt"
)
func main() {
	a:= 10
	fmt.Println(a<<3)
	fmt.Println(a>>3)
		}
output:
80
1
		
		
Program7
	
package main
import (
	"fmt"
)
func main() {
	var n float64 = 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)
		}
output:	
2.1e+14, float64




Program8
		
package main
import (
	"fmt"
)
func main() {
	greeting := "hello"
	name := "Stacy"
	sayGreeting(greeting, name)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}
output:

hello Stacy
Ted