package useralerts

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertsResponse 
// Deprecated: This class is obsolete. Use alertsGetResponse instead.
type AlertsResponse struct {
    AlertsGetResponse
}
// NewAlertsResponse instantiates a new alertsResponse and sets the default values.
func NewAlertsResponse()(*AlertsResponse) {
    m := &AlertsResponse{
        AlertsGetResponse: *NewAlertsGetResponse(),
    }
    return m
}
// CreateAlertsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertsResponse(), nil
}
// AlertsResponseable 
// Deprecated: This class is obsolete. Use alertsGetResponse instead.
type AlertsResponseable interface {
    AlertsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
