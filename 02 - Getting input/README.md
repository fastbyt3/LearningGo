# Reading user input

### Command line Arguments

- can be done using the **os** package

##### Imports

- multiple imports can be done using:

```go
import (
	"os"
	"fmt"
)
```

- all args can be taken from : `os.Args` which returns an array
- the first element(`os.Args[0]`) will be the prog

### During runtime from cmdline

- one option is to use `fmt.Scanf()` which works similar to `scanf()` from C

```go
fmt.Scanf("%s", &foo)
```

- another option is to use `NewReader` from **bufio** package
- Buffered I/O offers better performance
- we create a "Reader" first (like in Java)
- then we read the user inp and assign it to a var

```go
reader := bufio.NewReader(os.Stdin)
str1, _ := reader.ReadString('\n')
```

- we need to use `_` which is a empty var(allows only read) cos ReadString returns two values(the text, error(if any))
- `bufio.NewReader.ReadString('\n')` : `\n` is the delimiter

---
