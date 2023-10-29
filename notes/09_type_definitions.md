# Type Definitions

## Types

### Aliases vs. Definitions

You can add methods to new types, not to aliases!

```go
// Type Alias
type distance = float32

// New Type Based on
type distance float32

// Types Have a Constructor and Conversion Functions
d := distance(34.5)
```

### Methods

```go
type distanceMiles float32
type distanceKm float32

// Method
func (miles distanceMiles) ToKm() distanceKm {
    return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distanceMiles {
    return distanceMiles(km / 1.60934)
}

func test() {
    d := distanceMiles(4.5)
    km := d.ToKm()
}
```

## Complex Types for Definitions

### Structures

- They kind of replace the class idea
- It's a data type with strongly typed properties
- They have a default constructor
- You can add methods to it

#### Definitions

```go
type User struct {
    id int
    name string
}

func main() {
    var u1 User
    u1 = User {id: 1, name: "Frontend Masters"}
    u2 := User {2, "Frontend Masters"}
}
```

#### Methods

```go
func (u User) PrettyPrint() string {
    return string(u.id) + ": " + u.name
}

func main() {
    u2 := User {2, "Frontend Masters"}
    msg := u2.PrettyPrint()
}
```

### Interfaces

- A definition of methods
- You can emulate polymorphism from OOP
- Implicit implementation
- We can embed interfaces in other interfaces
