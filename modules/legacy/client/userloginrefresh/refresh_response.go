package userloginrefresh

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RefreshResponse 
// Deprecated: This class is obsolete. Use refreshPostResponse instead.
type RefreshResponse struct {
    RefreshPostResponse
}
// NewRefreshResponse instantiates a new refreshResponse and sets the default values.
func NewRefreshResponse()(*RefreshResponse) {
    m := &RefreshResponse{
        RefreshPostResponse: *NewRefreshPostResponse(),
    }
    return m
}
// CreateRefreshResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRefreshResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRefreshResponse(), nil
}
// RefreshResponseable 
// Deprecated: This class is obsolete. Use refreshPostResponse instead.
type RefreshResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RefreshPostResponseable
}
