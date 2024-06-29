package contact_formlist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Contact_formListRequestBuilder builds and executes requests for operations under \contact_form.list
type Contact_formListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Contact_formListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Contact_formListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContact_formListRequestBuilderInternal instantiates a new Contact_formListRequestBuilder and sets the default values.
func NewContact_formListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formListRequestBuilder) {
    m := &Contact_formListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contact_form.list", pathParameters),
    }
    return m
}
// NewContact_formListRequestBuilder instantiates a new Contact_formListRequestBuilder and sets the default values.
func NewContact_formListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContact_formListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all the contact forms that the current user
// Deprecated: This method is obsolete. Use GetAsListGetResponse instead.
func (m *Contact_formListRequestBuilder) Get(ctx context.Context, requestConfiguration *Contact_formListRequestBuilderGetRequestConfiguration)(ListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListResponseable), nil
}
// GetAsListGetResponse lists all the contact forms that the current user
func (m *Contact_formListRequestBuilder) GetAsListGetResponse(ctx context.Context, requestConfiguration *Contact_formListRequestBuilderGetRequestConfiguration)(ListGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListGetResponseable), nil
}
// ToGetRequestInformation lists all the contact forms that the current user
func (m *Contact_formListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Contact_formListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *Contact_formListRequestBuilder) WithUrl(rawUrl string)(*Contact_formListRequestBuilder) {
    return NewContact_formListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
