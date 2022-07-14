package ee

import "fmt"

// Phantom error, just a mark
var Phantom = fmt.Errorf("phantom error")
