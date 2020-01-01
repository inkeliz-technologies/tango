# IncrementalCamera Demo

## What does it do?
It demonstrates how one can move the camera gradually, instead of instantaneously. 

For doing so, it created a green background. This way, you'll notice the moving.  

## What are important aspects of the code?
This code is key in this demo:

```go
tango.Mailbox.Dispatch(tango.CameraMessage{
    Axis:        tango.ZAxis,
    Value:       3, // so zooming out a lot
    Incremental: true,
    Duration:    time.Second * 5,
})
```