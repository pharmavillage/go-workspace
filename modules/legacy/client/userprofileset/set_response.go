package userprofileset

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetResponse 
// Deprecated: This class is obsolete. Use setPostResponse instead.
type SetResponse struct {
    SetPostResponse
}
// NewSetResponse instantiates a new setResponse and sets the default values.
func NewSetResponse()(*SetResponse) {
    m := &SetResponse{
        SetPostResponse: *NewSetPostResponse(),
    }
    return m
}
// CreateSetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetResponse(), nil
}
// SetResponseable 
// Deprecated: This class is obsolete. Use setPostResponse instead.
type SetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SetPostResponseable
}
