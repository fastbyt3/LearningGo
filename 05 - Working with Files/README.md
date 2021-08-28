# Files in Go

### Checking if a file exists
- **os** package will be used for working with Files
- to check if a file exists we use : 

```go
os.Stat("fileName")
```

### Reading a file
- to read a file we need to import **io/ioutil**
- and invoke `ioutil.ReadFile()` this returns the content in *Bytes* and the error
- after converting from Bytes : `string(contentBytes)` we get the text

#### Reading line by line

- we need to use **bufio** which we used to read inputs
- code:

```go
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func main() {
  // open file for reading
  // read line by line
  lines, err := readLines("read2.go")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }
  // print file contents
  for i, line := range lines {
    fmt.Println(i, line)
  }
}
```

### Writing to a file

```go
openedFile.WriteString("text")
```
- we can either open an existing file or create a new one

```go
os.Create("new.txt")
os.Open("existing.txt")
```

- there are several ways to write to a file some involve overwriting the file and other approaches append data to file
- `ioutil.FileWrite(<file>, <str>, <permission/mode>)` this is a straight-forward way of overwriting a file
- 
