# Example usage of GRPC + Protobuf
This example demonstrates the description of a service in the `protos/garage.proto` file. This service does one thing; returns a random Car.

The service interface has been implemented as a go app in `garage_server/main.go`.

A client has been implemented in ruby in `garage_client/app.rb`

## Setup
Make sure you have Go installed `brew install go` and Ruby installed `brew install ruby`. Once you have Ruby installed you will also need Ruby's bundler `gem install bundler`.

Fetch this repo using `go get -u github.com/aren55555/grpc-example`

Then do `cd $GOPATH/src/github.com/aren55555/grpc-example`

## Running
First you need to convert the prot to Go and Ruby source files. To do this, from the root dir of this project execute `script/protogen`

Once this has completed in one terminal kick off the Go server `cd garage_server && go run main.go`. This will start the Go server - it will listen for incoming requests.

In another terminal kick off the Ruby client `cd garage_client && bundle && ruby app.rb`. This will start the Ruby client - it will send off requests in an infinite loop.

You should see print outs in both windows indicating the Ruby client is communicating with the Go server.
