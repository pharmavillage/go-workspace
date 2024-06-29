package models
import (
    "errors"
)
// Defines the source of the message
type Message_source int

const (
    EMAIL_MESSAGE_SOURCE Message_source = iota
    CHAT_MESSAGE_SOURCE
)

func (i Message_source) String() string {
    return []string{"email", "chat"}[i]
}
func ParseMessage_source(v string) (any, error) {
    result := EMAIL_MESSAGE_SOURCE
    switch v {
        case "email":
            result = EMAIL_MESSAGE_SOURCE
        case "chat":
            result = CHAT_MESSAGE_SOURCE
        default:
            return 0, errors.New("Unknown Message_source value: " + v)
    }
    return &result, nil
}
func SerializeMessage_source(values []Message_source) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Message_source) isMultiValue() bool {
    return false
}
