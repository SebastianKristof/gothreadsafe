# gothreadsafe

I wrote these Go containers to explore which operations need to run under a single lock. The useful case is a small service with shared in-memory state, such as a worker registry where several goroutines might try to register the same name.

## Try it

Run this from a Go 1.22.3+ module.

~~~sh
go get github.com/SebastianKristof/gothreadsafe/safemap
~~~

~~~go
package main

import (
    "fmt"

    "github.com/SebastianKristof/gothreadsafe/safemap"
)

func main() {
    workers := safemap.NewSafeMap[string, int]()
    workers.SetNX("transcriber", 1)
    workers.SetNX("transcriber", 2)

    if result := workers.Get("transcriber"); result.Found {
        fmt.Println(result.Value) // 1
    }
}
~~~

`SetNX` holds the lock for the entire check-and-insert operation. If you compose a separate read and write, another goroutine can change the map between those calls. The repository also includes [a slice wrapper](safeslice/) and [a set wrapper](safeset/) to explore the same problem across different collections.

The wrappers coordinate operations on the collection. If callers share a mutable value stored in a collection, they still need to coordinate access to that value.

## Run the checks

~~~sh
go test -race ./...
~~~

[The Go Tests workflow](.github/workflows/unit-tests.yaml) runs the suite with the Go race detector on pull requests.
