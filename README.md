# servo

A toy DSL and framework for cross-language API generation

## Servo DSL

### Options

Plugin options can be specified for configuring code generation.

```protobuf
option gojson.package = "package";
```

### Message

Messages have fields with specified types

```protobuf
message AllTypes {
  // message fields
  message_field: OtherType;
  // 32-bit integers
  int_field: int32;
  // 64-bit integers
  long_field: int64;
  // 32-bit floating points
  float_field: float32;
  // 64-bit floating points
  double_field: float64;
  // maps (key must be primitive)
  map_field: map[string]int;
  // lists
  list_field: list[string];
}
```

### RPC Service

Services can define RPCs that send and receive messages.

```protobuf
message EchoRequest {
  message: string;
}

message EchoResponse {
  message: string;
}

// Echo a message
service EchoService {
  rpc echo(EchoRequest): EchoResponse;
}
```

### PubSub Service

Services can define Publish methods that send messages
without responses.

```protobuf
message StatusUpdate {
  message: string;
}

service StatusService {
  pub update(StatusUpdate);
}
```