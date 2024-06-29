package usernotificationsreport

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportResponse 
// Deprecated: This class is obsolete. Use reportPostResponse instead.
type ReportResponse struct {
    ReportPostResponse
}
// NewReportResponse instantiates a new reportResponse and sets the default values.
func NewReportResponse()(*ReportResponse) {
    m := &ReportResponse{
        ReportPostResponse: *NewReportPostResponse(),
    }
    return m
}
// CreateReportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportResponse(), nil
}
// ReportResponseable 
// Deprecated: This class is obsolete. Use reportPostResponse instead.
type ReportResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReportPostResponseable
}
