# Visibility

What we write in a module:

- If it's camelCase, it's prvate
- If it's PascalCase, it's public

Variables and lambda functions can be:

- Module scoped
- Functions scoped
- Block scoped

```go
package main

import "fmt"

func notExported() {}   // private to the package

func Exported() {}     // public and exported to other packages
```