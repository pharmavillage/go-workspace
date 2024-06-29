package usernotificationsreport

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportPostRequestBody 
type ReportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The report_text property
    report_text *string
    // The reporter_email property
    reporter_email *string
    // The reporter_name property
    reporter_name *string
}
// NewReportPostRequestBody instantiates a new reportPostRequestBody and sets the default values.
func NewReportPostRequestBody()(*ReportPostRequestBody) {
    m := &ReportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["report_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportText(val)
        }
        return nil
    }
    res["reporter_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReporterEmail(val)
        }
        return nil
    }
    res["reporter_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReporterName(val)
        }
        return nil
    }
    return res
}
// GetReporterEmail gets the reporter_email property value. The reporter_email property
func (m *ReportPostRequestBody) GetReporterEmail()(*string) {
    return m.reporter_email
}
// GetReporterName gets the reporter_name property value. The reporter_name property
func (m *ReportPostRequestBody) GetReporterName()(*string) {
    return m.reporter_name
}
// GetReportText gets the report_text property value. The report_text property
func (m *ReportPostRequestBody) GetReportText()(*string) {
    return m.report_text
}
// Serialize serializes information the current object
func (m *ReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reporter_email", m.GetReporterEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reporter_name", m.GetReporterName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("report_text", m.GetReportText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReporterEmail sets the reporter_email property value. The reporter_email property
func (m *ReportPostRequestBody) SetReporterEmail(value *string)() {
    m.reporter_email = value
}
// SetReporterName sets the reporter_name property value. The reporter_name property
func (m *ReportPostRequestBody) SetReporterName(value *string)() {
    m.reporter_name = value
}
// SetReportText sets the report_text property value. The report_text property
func (m *ReportPostRequestBody) SetReportText(value *string)() {
    m.report_text = value
}
// ReportPostRequestBodyable 
type ReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReporterEmail()(*string)
    GetReporterName()(*string)
    GetReportText()(*string)
    SetReporterEmail(value *string)()
    SetReporterName(value *string)()
    SetReportText(value *string)()
}
