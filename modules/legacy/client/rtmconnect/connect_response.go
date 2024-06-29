package rtmconnect

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectResponse 
// Deprecated: This class is obsolete. Use connectGetResponse instead.
type ConnectResponse struct {
    ConnectGetResponse
}
// NewConnectResponse instantiates a new connectResponse and sets the default values.
func NewConnectResponse()(*ConnectResponse) {
    m := &ConnectResponse{
        ConnectGetResponse: *NewConnectGetResponse(),
    }
    return m
}
// CreateConnectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectResponse(), nil
}
// ConnectResponseable 
// Deprecated: This class is obsolete. Use connectGetResponse instead.
type ConnectResponseable interface {
    ConnectGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
