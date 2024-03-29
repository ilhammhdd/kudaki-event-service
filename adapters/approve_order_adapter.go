package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type ApproveOwnerOrder struct{}

func (ao *ApproveOwnerOrder) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ApproveOwnerOrder
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ao *ApproveOwnerOrder) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OwnerOrderApproved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
