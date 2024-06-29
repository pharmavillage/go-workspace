package actioninfo

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InfoResponse 
// Deprecated: This class is obsolete. Use infoGetResponse instead.
type InfoResponse struct {
    InfoGetResponse
}
// NewInfoResponse instantiates a new infoResponse and sets the default values.
func NewInfoResponse()(*InfoResponse) {
    m := &InfoResponse{
        InfoGetResponse: *NewInfoGetResponse(),
    }
    return m
}
// CreateInfoResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInfoResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInfoResponse(), nil
}
// InfoResponseable 
// Deprecated: This class is obsolete. Use infoGetResponse instead.
type InfoResponseable interface {
    InfoGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
