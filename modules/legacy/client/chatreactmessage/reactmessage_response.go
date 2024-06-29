package chatreactmessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReactmessageResponse 
// Deprecated: This class is obsolete. Use reactmessagePostResponse instead.
type ReactmessageResponse struct {
    ReactmessagePostResponse
}
// NewReactmessageResponse instantiates a new reactmessageResponse and sets the default values.
func NewReactmessageResponse()(*ReactmessageResponse) {
    m := &ReactmessageResponse{
        ReactmessagePostResponse: *NewReactmessagePostResponse(),
    }
    return m
}
// CreateReactmessageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReactmessageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReactmessageResponse(), nil
}
// ReactmessageResponseable 
// Deprecated: This class is obsolete. Use reactmessagePostResponse instead.
type ReactmessageResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReactmessagePostResponseable
}
