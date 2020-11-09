Program1

package main

import (
	"fmt"
)
func main() {
grade1 := 97
grade2 := 85
grade3 := 93
fmt.Printf("Grades %v, %v, %v", grade1, grade2, grade3)
}
output:
Grades 97, 85, 93	




Program2

package main
import (
	"fmt"
)
func main() {
grades := [3]int{97, 85, 93}

fmt.Printf("Grades %v", grades)
}
output:
Grades [97 85 93]

	


	
Program3	
	
package main
import (
	"fmt"
)
func main() {
var students  [3]string
fmt.Printf("students %v\n", students)
students[0] = "prathyusha"
fmt.Printf("students %v", students)
}
output:
students [  ]
students [prathyusha  ]




Program4

package main
import (
	"fmt"
)
func main() {
var students  [3]string
fmt.Printf("students %v\n", students)
students[0] = "prathyusha"
students[1] = "supraja"
students[2] = "damu"
fmt.Printf("students %v", students)
}
output:
students [  ]
students [prathyusha supraja damu]

	


Program5	
	
package main
import (
	"fmt"
)
func main() {
var students  [3]string
fmt.Printf("students %v\n", students)
students[0] = "prathyusha"
students[1] = "supraja"
students[2] = "damu"
fmt.Printf("students #1: %v", students[1])
}
output:
students [  ]
students #1: supraja



Program6

package main
import (
	"fmt"
)
func main() {
var students  [3]string
fmt.Printf("students %v\n", students)
students[0] = "prathyusha"
students[1] = "supraja"
students[2] = "damu"
fmt.Printf("students #1: %v\n", students[1])
fmt.Printf("number ofstudents %v", len(students))
}
output:
students [  ]
students #1: supraja
number ofstudents 3	




Program7

package main
import (
	"fmt"
)
func main() {
var identityMatrix  [3][3]int = [3][3] int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
fmt.Println(identityMatrix)
}
output:
[[1 0 0] [0 1 0] [0 0 1]]	
		