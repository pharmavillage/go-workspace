package fileupload

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadResponse 
// Deprecated: This class is obsolete. Use uploadPostResponse instead.
type UploadResponse struct {
    UploadPostResponse
}
// NewUploadResponse instantiates a new uploadResponse and sets the default values.
func NewUploadResponse()(*UploadResponse) {
    m := &UploadResponse{
        UploadPostResponse: *NewUploadPostResponse(),
    }
    return m
}
// CreateUploadResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUploadResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUploadResponse(), nil
}
// UploadResponseable 
// Deprecated: This class is obsolete. Use uploadPostResponse instead.
type UploadResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UploadPostResponseable
}
