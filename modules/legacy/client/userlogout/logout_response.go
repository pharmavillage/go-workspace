package userlogout

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LogoutResponse 
// Deprecated: This class is obsolete. Use logoutPostResponse instead.
type LogoutResponse struct {
    LogoutPostResponse
}
// NewLogoutResponse instantiates a new logoutResponse and sets the default values.
func NewLogoutResponse()(*LogoutResponse) {
    m := &LogoutResponse{
        LogoutPostResponse: *NewLogoutPostResponse(),
    }
    return m
}
// CreateLogoutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogoutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLogoutResponse(), nil
}
// LogoutResponseable 
// Deprecated: This class is obsolete. Use logoutPostResponse instead.
type LogoutResponseable interface {
    LogoutPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
