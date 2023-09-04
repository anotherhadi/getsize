// https://github.com/anotherhadi/getsize
package getsize

import (
	"golang.org/x/term"
	"os"
)

func GetSize() (width, height int, err error) {
	var fd uintptr = os.Stdin.Fd()
	width, height, err = term.GetSize(int(fd))
	return
}
