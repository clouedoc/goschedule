<p align="center">
# goschedule

Deadly Efficient Task System
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
import "github.com/jesuiscamille/goschedule
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

Please see LICENSE.md for license informations.

## TODO

- [ ] Documentation on exposed API functions
- [ ] Adding custom arguments to task functions
- [ ] Choosing precision
- [ ] Cool logo
- [X] Being cool
