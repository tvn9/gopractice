package main

import (
	"fmt"
	"os"
)

func main() {

	// Playing around with os package variables
	errs := []error{
		os.ErrInvalid,
		os.ErrPermission,
		os.ErrNotExist,
		os.ErrExist,
		os.ErrClosed,
		os.ErrNoDeadline,
		os.ErrDeadlineExceeded,
		os.ErrProcessDone,
	}
	fmt.Println(os.ErrInvalid)

	for i, e := range errs {
		fmt.Printf("#%d: %s\n", i, e)
	}

}
