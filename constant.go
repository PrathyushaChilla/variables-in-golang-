Program1

package main
import (
	"fmt"
)
const a int16 = 27
func main() {
	const a int = 14
	
	fmt.Printf("%v, %T\n", a, a)
		}
output:
14, int		
	
	


Program2

package main
import (
	"fmt"
)
const a int16 = 27
func main() {
	//const a int = 14
	
	fmt.Printf("%v, %T\n", a, a)
		}
output:
27, int16
		


		
Program3		
		
package main
import (
	"fmt"
)
func main() {
	const a int = 14
	const b int = 16
	fmt.Printf("%v, %T\n", a+b, a+b)
		}
outpuut:		
30, int





Program4
		
package main
import (
	"fmt"
)
const (
a = iota
b
c
)
func main() {
		
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
		}
output:
0
1
2





Program5

package main

import (
	"fmt"
)
const (
a = iota
b
c
)
const(
a2 = iota
)
func main() {
		
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
	
		}
output:
0
1
2
0





Program6
	
package main

import (
	"fmt"
)
const (
catSpecialist = iota
dogSpecialist 
snakeSpecialist
)
const(
a2 = iota
)
func main() {
var SpecialistType int = catSpecialist
		
	fmt.Printf("%v\n", SpecialistType == catSpecialist)
output:
true		
		

		
		
			
Program7

package main
import "fmt"
func main() {
    type Myint int
    const i1 Myint = 1
    const i2 = Myint(2)
    fmt.Printf("%T %v\n", i1, i1)
    fmt.Printf("%T %v\n", i2, i2)
}
output:
main.Myint 1
main.Myint 2
