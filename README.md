# go-desk
A Go CLI to interface with standing desks

My standing desk has a bluetooth interface. Let’s see if we can move it with Go


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
