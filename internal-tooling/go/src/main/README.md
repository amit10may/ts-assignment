# Instructions for Problem a
## Pre-requisites - Go Installed (v1.10+)
## Steps -
- Switch to directory go/src/main
- Execute - go run server.go
- This should start server on port 8080 (to receive json payloads) and 8081 (listing requests)
- Open another terminal and execute below commands
- telnet localhost 8080
- type the json payloads on single line. continue this for additional json requests.
- once all the jsons are sent, terminate the telnet session
- To see listing use below command
- telnet localhost 8081
- type count and hit enter (any text would do for now)
- This should print out the previously sent json request on separate lines

### TODO: The ports are hardcoded - can be made command line args
### Pending: The aspect of writing the json content to disk/database - Not implemented yet
