# Load environment variables from .env

## Installation:

```
go get github.com/prongbang/goenv
```

## Usage:

### `.env` file
```
DB_USER=root
DB_PASS=1234
```

### `main.go`
```golang
import (
	"fmt"
	"os"
	_ "github.com/prongbang/goenv"
)

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	fmt.Println("user:", user)
	fmt.Println("pass:", pass)
}
```

### Get ENV from file

.testenv

```
DB_USER=root
DB_PASS=1234
```

Used

```golang
import (
	"fmt"
	"os"
	"github.com/prongbang/goenv"
)

func main() {
	err := goenv.LoadEnv(".testenv")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
}
```
