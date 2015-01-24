# argf

ARGF for Golang

## Usage

### Use all arguments as file

```go
import "github.com/yuya-takeyama/argf"

func main() {
	reader, err := argf.Argf()
}
```

### With library like "flag"

```go
import (
	"flag"
	"github.com/yuya-takeyama/argf"
)

func main() {
	var displayHelp = flag.Bool("h", false, "display help")
	flag.Parse()

	reader, err := argf.From(flag.Args())
}
```

## Author

Yuya Takeyama
