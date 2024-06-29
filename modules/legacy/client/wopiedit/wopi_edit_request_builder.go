package wopiedit

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WopiEditRequestBuilder builds and executes requests for operations under \wopi.edit
type WopiEditRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WopiEditRequestBuilderGetQueryParameters this call is edit supported office document. The documents supported are docx, xlsx, pptx. doc, xls, ppt will result in conversion and saving a new file. A valid office 365 account is required to edit document.The response will be a HTML page to be shown in a new window with embedded script to begin the editing process.
type WopiEditRequestBuilderGetQueryParameters struct {
    // path to location
    Fspath *string `uriparametername:"fspath"`
}
// WopiEditRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WopiEditRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WopiEditRequestBuilderGetQueryParameters
}
// NewWopiEditRequestBuilderInternal instantiates a new WopiEditRequestBuilder and sets the default values.
func NewWopiEditRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WopiEditRequestBuilder) {
    m := &WopiEditRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/wopi.edit?fspath={fspath}", pathParameters),
    }
    return m
}
// NewWopiEditRequestBuilder instantiates a new WopiEditRequestBuilder and sets the default values.
func NewWopiEditRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WopiEditRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWopiEditRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this call is edit supported office document. The documents supported are docx, xlsx, pptx. doc, xls, ppt will result in conversion and saving a new file. A valid office 365 account is required to edit document.The response will be a HTML page to be shown in a new window with embedded script to begin the editing process.
func (m *WopiEditRequestBuilder) Get(ctx context.Context, requestConfiguration *WopiEditRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateEdit400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation this call is edit supported office document. The documents supported are docx, xlsx, pptx. doc, xls, ppt will result in conversion and saving a new file. A valid office 365 account is required to edit document.The response will be a HTML page to be shown in a new window with embedded script to begin the editing process.
func (m *WopiEditRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WopiEditRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/html, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WopiEditRequestBuilder) WithUrl(rawUrl string)(*WopiEditRequestBuilder) {
    return NewWopiEditRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
