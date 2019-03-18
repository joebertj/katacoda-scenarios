All programming tutorials begin with Hello World

## Task

Create a file using `vi`

`vi hello-world.go`{{execute}}

Type this: (Make sure you press **a** or **i** first to toggle edit mode in `vi`)

`i`{{execute}}

```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```{{execute}}

Save your work.(Make sure you press **ESC** first to toggle command mode in `vi`)

`\u001b[`{{execute}}

I have remaped **ESC** to **jj**

`jj`{{execute}}

`:wq`{{execute}}
