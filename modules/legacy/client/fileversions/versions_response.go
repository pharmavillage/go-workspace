package fileversions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VersionsResponse 
// Deprecated: This class is obsolete. Use versionsGetResponse instead.
type VersionsResponse struct {
    VersionsGetResponse
}
// NewVersionsResponse instantiates a new versionsResponse and sets the default values.
func NewVersionsResponse()(*VersionsResponse) {
    m := &VersionsResponse{
        VersionsGetResponse: *NewVersionsGetResponse(),
    }
    return m
}
// CreateVersionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVersionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVersionsResponse(), nil
}
// VersionsResponseable 
// Deprecated: This class is obsolete. Use versionsGetResponse instead.
type VersionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VersionsGetResponseable
}
