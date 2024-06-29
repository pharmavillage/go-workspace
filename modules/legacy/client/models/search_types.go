package models
import (
    "errors"
)
// 
type SearchTypes int

const (
    MESSAGE_SEARCHTYPES SearchTypes = iota
    ACTION_SEARCHTYPES
    FILE_SEARCHTYPES
    USER_SEARCHTYPES
)

func (i SearchTypes) String() string {
    return []string{"message", "action", "file", "user"}[i]
}
func ParseSearchTypes(v string) (any, error) {
    result := MESSAGE_SEARCHTYPES
    switch v {
        case "message":
            result = MESSAGE_SEARCHTYPES
        case "action":
            result = ACTION_SEARCHTYPES
        case "file":
            result = FILE_SEARCHTYPES
        case "user":
            result = USER_SEARCHTYPES
        default:
            return 0, errors.New("Unknown SearchTypes value: " + v)
    }
    return &result, nil
}
func SerializeSearchTypes(values []SearchTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SearchTypes) isMultiValue() bool {
    return false
}
