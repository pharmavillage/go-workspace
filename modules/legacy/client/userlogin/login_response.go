package userlogin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoginResponse 
// Deprecated: This class is obsolete. Use loginPostResponse instead.
type LoginResponse struct {
    LoginPostResponse
}
// NewLoginResponse instantiates a new loginResponse and sets the default values.
func NewLoginResponse()(*LoginResponse) {
    m := &LoginResponse{
        LoginPostResponse: *NewLoginPostResponse(),
    }
    return m
}
// CreateLoginResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoginResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoginResponse(), nil
}
// LoginResponseable 
// Deprecated: This class is obsolete. Use loginPostResponse instead.
type LoginResponseable interface {
    LoginPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
