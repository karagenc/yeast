# Yeast

This is the Go implementation of https://github.com/unshiftio/yeast.

Yeast is a unique ID generator. It has been primarily designed to generate a unique ID which can be used for cache busting. A common practice for this is to use a timestamp, but there are couple of downsides when using timestamps:

1. The timestamp is already 13 chars long. This might not matter for 1 request, but if you make hundreds of them, this quickly adds up in bandwidth and processing time.
2. It's not unique enough. If you generate two stamps right after each other, they would be identical because the timing accuracy is limited to milliseconds.

Yeast solves both of these issues by:

1. Compressing the generated timestamp using a custom encode() function that returns a string representation of the number.
2. Seeding the id in case of collision (when the id is identical to the previous one).

To keep the strings unique it will use the `.` char to separate the generated stamp from the seed.

## Usage

**Note:** All functions of the `Yeaster` can be concurrently accessed.

```go
import (
    "github.com/karagenc/yeast"
    "fmt"
    "time"
)

y := yeast.New()

// Yeast function generates unique ID's using the current time.
// It uses the Encode function under the hood.
//
// Generate 3 IDs:
id1 := y.Yeast()
id2 := y.Yeast()
id3 := y.Yeast()
fmt.Printf("IDs:\n%s\n%s\n%s\n", id1, id2, id3)

// Encode the given number
s := y.Encode(1337)
// Decode the string as number
n, _ := y.Decode(s)
fmt.Printf("n: %d\n", n) // Outputs 1337
```
