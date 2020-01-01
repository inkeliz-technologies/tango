# Scenes Demo

## What does it do?
It demonstrates how one can use multiple `Scene`s, and switch between them.  

## What are important aspects of the code?
These things are key in this demo:

* Defining two Scenes: `IconScene` and `RockScene`
* Giving one to `tango.Open` as the default `Scene`
* Registering the other with `tango.RegisterScene`, so we can later:
* Call `tango.SetSceneByName` to switch the `Scene`s. 
