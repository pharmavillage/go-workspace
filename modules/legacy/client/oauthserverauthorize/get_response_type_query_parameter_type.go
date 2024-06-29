package oauthserverauthorize
import (
    "errors"
)
// 
type GetResponse_typeQueryParameterType int

const (
    CODE_GETRESPONSE_TYPEQUERYPARAMETERTYPE GetResponse_typeQueryParameterType = iota
)

func (i GetResponse_typeQueryParameterType) String() string {
    return []string{"code"}[i]
}
func ParseGetResponse_typeQueryParameterType(v string) (any, error) {
    result := CODE_GETRESPONSE_TYPEQUERYPARAMETERTYPE
    switch v {
        case "code":
            result = CODE_GETRESPONSE_TYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetResponse_typeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetResponse_typeQueryParameterType(values []GetResponse_typeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_typeQueryParameterType) isMultiValue() bool {
    return false
}
