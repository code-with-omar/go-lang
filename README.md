## Data types

### Basic Data types

### 1.Numeric Types

- Integer
  - `Signed`: int, int8, int16, int32, int64
  - `Unsigned`: uint, uint8 (alias byte), uint16, uint32, uint64, uintptr
- `Floating-point`: float32, float64
- `Complex numbers`: complex64, complex128
- `Boolean`
  bool (values: true or false)
- String
  - string (immutable sequence of bytes; UTF-8 encoded)

## Variable Declarations

### 1 `var` with type Inference

```golang
var city = "Comilla" // Inference as string
var temparature = 30.5 // Inference as float64
```

### 2. Short Declaration (`:=`)

```golang
message := "Hello, GO!"
count:=10
```

- Note: ⚠️ := can only be used inside functions.

### 3. Declaring Multiple Variables

```golang
var x,y,z int
var p,q,r int =5,6,7
var (
    title string = "Developer"
    salary float64 = 5000.00
    active bool = true
)
```

### 4. Zero Values

```golang
var a int
var b string
var c bool
```
