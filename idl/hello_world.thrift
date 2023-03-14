include "base.thrift"

namespace go tangerine.go.thrift.hello_world

struct SayRequest {
    255: optional base.RPCRequest Base
}

struct SayResponse {
    1: required string content
    255: required base.RPCResponse Base
}

service HelloWorldSayer {
    SayResponse say(1:SayRequest req)
}