# javastyle-collection

## PriorityQueue
```go
package main

import (
	"fmt"
	"github.com/nsce9806q/javastyle-collection/priorityqueue"
)

func main() {
	pq := priorityqueue.New[int]()

	pq.Add(3)
	pq.Add(2)
	pq.Add(1)

	fmt.Println(pq.Poll()) // 1
	fmt.Println(pq.Poll()) // 2
	fmt.Println(pq.Poll()) // 3
}
```

```go
package main

import (
	"fmt"
	"github.com/nsce9806q/javastyle-collection/priorityqueue"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	pq := priorityqueue.New(priorityqueue.WithComparator(func(a, b Person) int {
		return a.Age - b.Age
	}))

	p1 := Person{Name: "민수", Age: 28}
	p2 := Person{Name: "찬우", Age: 29}
	p3 := Person{Name: "기범", Age: 25}

	pq.Add(p1)
	pq.Add(p2)
	fmt.Println(pq.Contains(p3)) // false

	pq.Add(p3)
	fmt.Println(pq.Contains(p3)) // true

	pqArray := pq.ToArray()
	fmt.Println(pqArray) // [{Name: 기범, Age: 25} {Name: 찬우, Age: 29} {Name: 민수, Age: 28}]

	fmt.Println(pq.Size()) // 3
	fmt.Println(pq.Peek()) // {Name: 기범, Age: 25}

	fmt.Println(pq.Poll()) // {Name: 기범, Age: 25}
	fmt.Println(pq.Poll()) // {Name: 민수, Age: 28}
	fmt.Println(pq.Poll()) // {Name: 찬우, Age: 29}
}
```