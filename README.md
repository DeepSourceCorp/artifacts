# Contracts

Contract definitions for services in DeepSource.

### Error codes

| Level   | Code |
|---------|------|
| ERROR   | 0    |
| WARNING | 1    |

### Running tests
 - Install [uber/prototool](https://github.com/uber/prototool) and make sure it's available in your path.
 - Run `make test`.


### Beacon

- To generate golang stubs

```
brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc -I=./ --go_out=plugins=grpc:./ ./beacon.proto
```
