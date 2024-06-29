package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchResult 
type SearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ocurrences of the search highlighted on contex.
    highlighted *string
    // The subject property
    subject SearchResult_SearchResult_subjectable
    // The type property
    typeEscaped *SearchTypes
}
// SearchResult_SearchResult_subject composed type wrapper for classes ActionSearchResult, FileSearchResult, MessageSearchResult, UserSearchResult
type SearchResult_SearchResult_subject struct {
    // Composed type representation for type ActionSearchResult
    actionSearchResult ActionSearchResultable
    // Composed type representation for type FileSearchResult
    fileSearchResult FileSearchResultable
    // Composed type representation for type MessageSearchResult
    messageSearchResult MessageSearchResultable
    // Composed type representation for type UserSearchResult
    userSearchResult UserSearchResultable
}
// NewSearchResult_SearchResult_subject instantiates a new SearchResult_subject and sets the default values.
func NewSearchResult_SearchResult_subject()(*SearchResult_SearchResult_subject) {
    m := &SearchResult_SearchResult_subject{
    }
    return m
}
// CreateSearchResult_SearchResult_subjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchResult_SearchResult_subjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSearchResult_SearchResult_subject()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetActionSearchResult gets the ActionSearchResult property value. Composed type representation for type ActionSearchResult
func (m *SearchResult_SearchResult_subject) GetActionSearchResult()(ActionSearchResultable) {
    return m.actionSearchResult
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResult_SearchResult_subject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFileSearchResult gets the FileSearchResult property value. Composed type representation for type FileSearchResult
func (m *SearchResult_SearchResult_subject) GetFileSearchResult()(FileSearchResultable) {
    return m.fileSearchResult
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *SearchResult_SearchResult_subject) GetIsComposedType()(bool) {
    return true
}
// GetMessageSearchResult gets the MessageSearchResult property value. Composed type representation for type MessageSearchResult
func (m *SearchResult_SearchResult_subject) GetMessageSearchResult()(MessageSearchResultable) {
    return m.messageSearchResult
}
// GetUserSearchResult gets the UserSearchResult property value. Composed type representation for type UserSearchResult
func (m *SearchResult_SearchResult_subject) GetUserSearchResult()(UserSearchResultable) {
    return m.userSearchResult
}
// Serialize serializes information the current object
func (m *SearchResult_SearchResult_subject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionSearchResult() != nil {
        err := writer.WriteObjectValue("", m.GetActionSearchResult())
        if err != nil {
            return err
        }
    } else if m.GetFileSearchResult() != nil {
        err := writer.WriteObjectValue("", m.GetFileSearchResult())
        if err != nil {
            return err
        }
    } else if m.GetMessageSearchResult() != nil {
        err := writer.WriteObjectValue("", m.GetMessageSearchResult())
        if err != nil {
            return err
        }
    } else if m.GetUserSearchResult() != nil {
        err := writer.WriteObjectValue("", m.GetUserSearchResult())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionSearchResult sets the ActionSearchResult property value. Composed type representation for type ActionSearchResult
func (m *SearchResult_SearchResult_subject) SetActionSearchResult(value ActionSearchResultable)() {
    m.actionSearchResult = value
}
// SetFileSearchResult sets the FileSearchResult property value. Composed type representation for type FileSearchResult
func (m *SearchResult_SearchResult_subject) SetFileSearchResult(value FileSearchResultable)() {
    m.fileSearchResult = value
}
// SetMessageSearchResult sets the MessageSearchResult property value. Composed type representation for type MessageSearchResult
func (m *SearchResult_SearchResult_subject) SetMessageSearchResult(value MessageSearchResultable)() {
    m.messageSearchResult = value
}
// SetUserSearchResult sets the UserSearchResult property value. Composed type representation for type UserSearchResult
func (m *SearchResult_SearchResult_subject) SetUserSearchResult(value UserSearchResultable)() {
    m.userSearchResult = value
}
// SearchResult_SearchResult_subjectable 
type SearchResult_SearchResult_subjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSearchResult()(ActionSearchResultable)
    GetFileSearchResult()(FileSearchResultable)
    GetMessageSearchResult()(MessageSearchResultable)
    GetUserSearchResult()(UserSearchResultable)
    SetActionSearchResult(value ActionSearchResultable)()
    SetFileSearchResult(value FileSearchResultable)()
    SetMessageSearchResult(value MessageSearchResultable)()
    SetUserSearchResult(value UserSearchResultable)()
}
// NewSearchResult instantiates a new SearchResult and sets the default values.
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["highlighted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlighted(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchResult_SearchResult_subjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(SearchResult_SearchResult_subjectable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SearchTypes))
        }
        return nil
    }
    return res
}
// GetHighlighted gets the highlighted property value. The ocurrences of the search highlighted on contex.
func (m *SearchResult) GetHighlighted()(*string) {
    return m.highlighted
}
// GetSubject gets the subject property value. The subject property
func (m *SearchResult) GetSubject()(SearchResult_SearchResult_subjectable) {
    return m.subject
}
// GetTypeEscaped gets the type property value. The type property
func (m *SearchResult) GetTypeEscaped()(*SearchTypes) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("highlighted", m.GetHighlighted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *SearchResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHighlighted sets the highlighted property value. The ocurrences of the search highlighted on contex.
func (m *SearchResult) SetHighlighted(value *string)() {
    m.highlighted = value
}
// SetSubject sets the subject property value. The subject property
func (m *SearchResult) SetSubject(value SearchResult_SearchResult_subjectable)() {
    m.subject = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *SearchResult) SetTypeEscaped(value *SearchTypes)() {
    m.typeEscaped = value
}
// SearchResultable 
type SearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHighlighted()(*string)
    GetSubject()(SearchResult_SearchResult_subjectable)
    GetTypeEscaped()(*SearchTypes)
    SetHighlighted(value *string)()
    SetSubject(value SearchResult_SearchResult_subjectable)()
    SetTypeEscaped(value *SearchTypes)()
}
