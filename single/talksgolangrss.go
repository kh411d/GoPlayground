package main //https://talks.golang.org/2013/advconc.slide#11
import (
    "fmt"
    "time"
)

type Fetcher interface {
    Fetch(uri string) (items []item, next time.Time, err Error)
}

type item struct {
    Title, Channel, GUID string
}

func Fetch(domain string) Fetcher {

}

type Subscription interface {
    Updates() <-chan item
    Close() error
}

// sub implements the Subscription interface.
type sub struct {
    fetcher Fetcher
    updates chan item
}

// loop fetches items using s.fetcher and sends them
// on s.updates.  loop exits when s.Close is called.
func (s *sub) loop() {

}

func (s *sub) Updates() <-chan Item {
    return s.updates
}

func (s *sub) Close() error {
    // TODO: make loop exit
    // TODO: find out about any error
    return err
}

//convert fetches to stream
func Subscribe(fetcher Fetcher) Subscription {
    s := &sub{
        fetcher: fetcher,
        updates: make(chan Item), // for Updates
    }
    go s.loop()
    return s
}

func Merge(subs ...Subscription) Subscription {

}

func main() {
    // Subscribe to some feeds, and create a merged update stream.
    merged := Merge(
        Subscribe(Fetch("blog.golang.org")),
        Subscribe(Fetch("googleblog.blogspot.com")),
        Subscribe(Fetch("googledevelopers.blogspot.com")))

    // Close the subscriptions after some time.
    time.AfterFunc(3*time.Second, func() {
        fmt.Println("closed:", merged.Close())
    })

    // Print the stream.
    for it := range merged.Updates() {
        fmt.Println(it.Channel, it.Title)
    }

    panic("show me the stacks")
}
