<p align="center">
<h1 align="center"><b>goschedule</b></h1>
<h4 align="center">~ A Deadly Efficient Tasking System for Golang ~</h4>
</p>

## Features
* safe for concurent use
    - use it accross goroutines !
* precise to the minute
* lightweight
    - efficient memory management
* intuitive API
* easy task management
    - add tasks
    - remove tasks

## Installation

```
go get github.com/jesuiscamille/goschedule
```

To use it in your project, add this:

```go
import "github.com/jesuiscamille/goschedule"
```

## Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/jesuiscamille/goschedule"
)

func main() {
    // This creates a scheduler.
    scheduler := goschedule.NewScheduler()

    // Adding a task will automagically start the scheduler
    id := scheduler.AddTask(15, 30, func() {
        now := time.Now()
        fmt.Printf("Hello, it's %d:%d !\n", now.Hour, now.Minute)
    })

    _ = id

    for {
        time.Sleep(time.Hour)
    }
}
```

If you change your mind, you can remove a task:

```go
// If no task remains, the scheduler will automagically
// be stopped.
// Note: id is obtained as a result of scheduler.AddTask(...)
_ = scheduler.RemoveTask(id)
```

## LICENSE

This code is licensed under the GNU GPLv3.
*grosso-modo*, you can do whatever you want with the code *albeit* using it in closed source commercial projects without sharing back your modifications.
Please see LICENSE for more details.

## TODO

- [ ] Documentation
- [ ] Custom Arguments
- [ ] Precision
- [ ] Cool logo
- [X] Being cool

I usually remove items I'm working on from the list, so feel free to make your changes and open a pull request !
