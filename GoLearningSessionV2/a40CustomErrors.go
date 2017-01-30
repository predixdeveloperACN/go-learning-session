package main

import (
    "fmt"
    "time"
    "runtime"
    "log"
    "errors"
)

// Go lets you can make your own custom Error objects.
//
type MyError struct {
    When time.Time
    What string
}

// We can print our errors by implementing the Error function.
//
func (e MyError) Error() string {
    return fmt.Sprintf("At %v, %s", e.When, e.What)
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
    // SAMPLE OUTPUT:
    // At 2017-01-05 10:33:10.6909022 +0800 CST, an error occured.

    // ---------------------------------------------
    //
    // You can also check for a specific kind of error and behave accordingly.
    //
    if err := run(); err != nil {
        switch err.(type) {
        case MyError:
            fmt.Println("Custom error: " + err.Error())
        default:
            fmt.Println("Unknown error: " + err.Error())
        }
    }
    // SAMPLE OUTPUT:
    // Custom error: At 2017-01-17 08:31:21.6257128 +0800 CST, an error occured.

    // ---------------------------------------------
    //
    // The following code demonstrates 2 functions that
    // log errors and print the line number, filepath, etc.
    //
    // Note that for Panics, all the information you will most likely need
    // (e.g. line number, filepath, etc.) are automatically included.

    if BasicErrorLog(errors.New("An error has occurred!")) {
        log.Print("doing other stuff")
    }
    // SAMPLE OUTPUT:
    // 2017/01/17 08:51:14 [error] D:/VinceGo/src/GoLearningSession/40CustomErrors.go:46 An error has occurred!
    // 2017/01/17 08:51:14 doing other stuff

    if FullErrorLog(errors.New("An error has occurred!")) {
        log.Print("doing other stuff")
    }
    // SAMPLE OUTPUT:
    // 2017/01/17 08:51:14 [error] in main.main[D:/VinceGo/src/GoLearningSession/40CustomErrors.go:39] An error has occurred!
    // 2017/01/17 08:51:14 doing other stuff
}

// This test function simulates an error by returning a MyError object.
//
func run() error {
    return MyError{
        time.Now(),
        "an error occured.",
    }
}

func BasicErrorLog(err error) (hasError bool) {
    if err != nil {
        // Notice that we're using 1, so it will actually log where
        // the error happened; 0 means this function, and we don't want that.
        _, fn, line, _ := runtime.Caller(1)
        log.Printf("[error] %s:%d %v", fn, line, err)
        hasError = true
    }
    return
}

// This logs the package name as well as the function name
//
func FullErrorLog(err error) (hasError bool) {
    if err != nil {
        pc, fn, line, _ := runtime.Caller(1)
        log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
        hasError = true
    }
    return
}
