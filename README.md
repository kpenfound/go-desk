# go-desk
A Go CLI to interface with standing desks

Hashicorp Eng-serv hackathon spring 2021

My standing desk has a bluetooth interface. Let’s see if we can move it with Go and a public API for fun interactive-ness.


## Idea

- public http api to submit desk heights to (POST https://www.example.com/ {"height":95})

- go service to connect the api to the cli. Have api run on a lambda and submit integer heights to sqs, while a local service pulls from sqs and runs the cli

- go cli to change the desk height. Can use the python cli in it’s place before this step


## Existing work

My desk is an Idasen from Ikea. There's a python cli that interfaces with this: https://pypi.org/project/idasen/

A go bluetooth interface: https://github.com/tinygo-org/bluetooth


## Other thoughts

Initially this will be focused on the Idasen desk. But with a common interface, this could be expanded to other desks as well.

When the cli is functional, I'd like to set a macro on my keyboard to move the desk to different profiles.

## Result

Ended up with a lambda + api gateway endpoint which receives a `{ "position" : "sit | stand" }`, and puts the message in a SQS queue.
The `go-desk listen` subcommand listens to that queue and submits the commands to the desk cli which communicates to the desk via bluetooth.

For some unknown reason, the desk's bluetooth interface is not visible to my macbook, however it was visible to my other devices.  I ended up
doing the `go-desk listen` on my thinkpad which was able to pair with the desk.

For the future, I intend to set up a macropad on my thinkpad to submit the different saved profiles to the desk.  I also want to reimplement
the python idasen cli in go to learn more about bluetooth.
