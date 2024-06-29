package actionlist
import (
    "errors"
)
// 
type GetSort_byQueryParameterType int

const (
    DEFAULTESCAPED_GETSORT_BYQUERYPARAMETERTYPE GetSort_byQueryParameterType = iota
    NAME_GETSORT_BYQUERYPARAMETERTYPE
    CHANNEL_GETSORT_BYQUERYPARAMETERTYPE
    DUE_DATE_GETSORT_BYQUERYPARAMETERTYPE
)

func (i GetSort_byQueryParameterType) String() string {
    return []string{"default", "name", "channel", "due_date"}[i]
}
func ParseGetSort_byQueryParameterType(v string) (any, error) {
    result := DEFAULTESCAPED_GETSORT_BYQUERYPARAMETERTYPE
    switch v {
        case "default":
            result = DEFAULTESCAPED_GETSORT_BYQUERYPARAMETERTYPE
        case "name":
            result = NAME_GETSORT_BYQUERYPARAMETERTYPE
        case "channel":
            result = CHANNEL_GETSORT_BYQUERYPARAMETERTYPE
        case "due_date":
            result = DUE_DATE_GETSORT_BYQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSort_byQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSort_byQueryParameterType(values []GetSort_byQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSort_byQueryParameterType) isMultiValue() bool {
    return false
}
