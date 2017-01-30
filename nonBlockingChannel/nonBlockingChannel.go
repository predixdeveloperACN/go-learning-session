
// https://github.com/hectane/go-nonblockingchan
//
// ------------------------------------------------------
//
// A special type that mimics the behavior of a channel but does not block when items are sent.
//
// +++++
//
// Features:
//
// Send items without ever worrying that the send will block
// Check how many items are waiting to be received
// Synchronized access to members - use it from any goroutine
//
// +++++
//
// Usage:
//
// To use the package, add the following import:
//
//     import "github.com/hectane/go-nonblockingchan"
//
// Use the New() function to create a new instance:
//
//     c := nbc.New()
//
// To send an item on the channel, use the Send field:
//
//     c.Send <- true
//
// Sending will always succeed immediately. The item will be added to an internal buffer until it is received:
//
//     v, ok := <-c.Recv
//     if ok {
//         // value was received
//     } else {
//         // channel was closed
//     }
//
// ------------------------------------------------------

// Non-blocking channel for Go.
package nbc

import (
    "container/list"
    "sync"
)

// Special type that mimics the behavior of a channel but does not block when
// items are sent. Items are stored internally until received. Closing the Send
// channel will cause the Recv channel to be closed after all items have been received.
//
type NonBlockingChan struct {
    mutex sync.Mutex
    Send  chan<- interface{}
    Recv  <-chan interface{}
    items *list.List
}

// Loop for buffering items between the Send and Recv channels until the Send channel is closed.
//
func (n *NonBlockingChan) run(send <-chan interface{}, recv chan<- interface{}) {
    for {
        if send == nil && n.items.Len() == 0 {
            close(recv)
            break
        }

        var (
            recvChan chan<- interface{}
            recvVal  interface{}
        )

        if n.items.Len() > 0 {
            recvChan = recv
            recvVal = n.items.Front().Value
        }

        select {

        case i, ok := <-send:
            if ok {
                n.mutex.Lock()
                n.items.PushBack(i)
                n.mutex.Unlock()
            } else {
                send = nil
            }

        case recvChan <- recvVal:
            n.mutex.Lock()
            n.items.Remove(n.items.Front())
            n.mutex.Unlock()
        }
    }
}

// Create a new non-blocking channel.
func New() *NonBlockingChan {
    var (
        send = make(chan interface{})
        recv = make(chan interface{})
        n    = &NonBlockingChan{
            Send:  send,
            Recv:  recv,
            items: list.New(),
        }
    )
    go n.run(send, recv)
    return n
}

// Retrieve the number of items waiting to be received.
func (n *NonBlockingChan) Len() int {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    return n.items.Len()
}
