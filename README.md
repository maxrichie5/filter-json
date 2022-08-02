# filter-json

A simple program that reads from stdin and filters all valid json lines to stdout and all other lines to stderr.

## Installing

To install, clone this repo and run 

```shell
go build && go install
```
## Usage

This program was made to be able to handle local logs from services easier. 
With this, you can pipe logs to `jq` to pretty print them. 
Logs are written to stderr in services, so they first need to be redirected to stdout to be used with a pipe. 
The full command will look like this 

```shell
make run 3>&1 1>&2 2>&3 | filter-json | jq . 
```

This run the program, swaps stdout and stderr, filters json lines, and then pretty-prints them with JQ.

It is recommended to combine this with an alias in your rc file to make it super easy to run the whole thing. For example:

```shell
alias run="make run 3>&1 1>&2 2>&3 | filter-json | jq .event"
```

## Supporting the Dev

This is an open-source project made by a poor, little dev named Max.
If you pity him and appreciate this project, you can give him all your monies. Thank you.
