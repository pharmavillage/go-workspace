package channelhistory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HistoryResponse 
// Deprecated: This class is obsolete. Use historyGetResponse instead.
type HistoryResponse struct {
    HistoryGetResponse
}
// NewHistoryResponse instantiates a new historyResponse and sets the default values.
func NewHistoryResponse()(*HistoryResponse) {
    m := &HistoryResponse{
        HistoryGetResponse: *NewHistoryGetResponse(),
    }
    return m
}
// CreateHistoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHistoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHistoryResponse(), nil
}
// HistoryResponseable 
// Deprecated: This class is obsolete. Use historyGetResponse instead.
type HistoryResponseable interface {
    HistoryGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
