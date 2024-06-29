package systeminfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SystemInfoRequestBuilder builds and executes requests for operations under \system.info
type SystemInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SystemInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SystemInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSystemInfoRequestBuilderInternal instantiates a new SystemInfoRequestBuilder and sets the default values.
func NewSystemInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemInfoRequestBuilder) {
    m := &SystemInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system.info", pathParameters),
    }
    return m
}
// NewSystemInfoRequestBuilder instantiates a new SystemInfoRequestBuilder and sets the default values.
func NewSystemInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSystemInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get system.info - Get information about the system
// Deprecated: This method is obsolete. Use GetAsInfoGetResponse instead.
func (m *SystemInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *SystemInfoRequestBuilderGetRequestConfiguration)(InfoResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoResponseable), nil
}
// GetAsInfoGetResponse system.info - Get information about the system
func (m *SystemInfoRequestBuilder) GetAsInfoGetResponse(ctx context.Context, requestConfiguration *SystemInfoRequestBuilderGetRequestConfiguration)(InfoGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoGetResponseable), nil
}
// ToGetRequestInformation system.info - Get information about the system
func (m *SystemInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SystemInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SystemInfoRequestBuilder) WithUrl(rawUrl string)(*SystemInfoRequestBuilder) {
    return NewSystemInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
