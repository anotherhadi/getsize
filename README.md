<p align="center">
</p>

# GetSize

GetSize is a simple Go package that allows you to retrieve the size (width and height) of the terminal window. This can be useful for creating text-based applications and ensuring proper formatting within the terminal.


---

## Usage

### Installation:
```bash
go get -u github.com/anotherhadi/getsize@latest
```

### Importing:
```go
import (
  "github.com/anotherhadi/getsize"
)
```

### Basic Usage:
```go
func main() {
  width, height, err := getsize.GetSize()
  if err != nil {
    panic(err)
  }
  fmt.Println("Width:", width, " Height:", height)
}
```

---

<img src="https://img.buymeacoffee.com/button-api/?text=Buy me a cookie&emoji=🍪&slug=anotherhadi&button_colour=eed2cc&font_colour=000000&font_family=Inter&outline_colour=ffffff&coffee_colour=ff0000" />

