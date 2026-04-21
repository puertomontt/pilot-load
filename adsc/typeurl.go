package adsc

import "google.golang.org/protobuf/proto"

func typeURL[T proto.Message]() string {
	ft := new(T)
	return "type.googleapis.com/" + string((*ft).ProtoReflect().Descriptor().FullName())
}
