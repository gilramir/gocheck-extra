# gocheck-extra

This Go module has types that are useful when writing unit tests with the
wonderful gopkg.in/check.v1 module.

Import it similarly to how you would import gopkg.in/check.v1:
```
import (
        . "github.com/gilramir/gocheck-extra"
        . "gopkg.in/check.v1"
)
```

## InSlice

This checker lets you check if an item is in a slice.
```
values := []int{2, 4, 8}
c.Check(2, InSlice, values)
```

