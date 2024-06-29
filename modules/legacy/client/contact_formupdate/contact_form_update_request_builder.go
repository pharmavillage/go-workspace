package contact_formupdate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Contact_formUpdateRequestBuilder builds and executes requests for operations under \contact_form.update
type Contact_formUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Contact_formUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Contact_formUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContact_formUpdateRequestBuilderInternal instantiates a new Contact_formUpdateRequestBuilder and sets the default values.
func NewContact_formUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formUpdateRequestBuilder) {
    m := &Contact_formUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contact_form.update", pathParameters),
    }
    return m
}
// NewContact_formUpdateRequestBuilder instantiates a new Contact_formUpdateRequestBuilder and sets the default values.
func NewContact_formUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContact_formUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates an existing contact form
// Deprecated: This method is obsolete. Use PostAsUpdatePostResponse instead.
func (m *Contact_formUpdateRequestBuilder) Post(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *Contact_formUpdateRequestBuilderPostRequestConfiguration)(UpdateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdateResponseable), nil
}
// PostAsUpdatePostResponse updates an existing contact form
func (m *Contact_formUpdateRequestBuilder) PostAsUpdatePostResponse(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *Contact_formUpdateRequestBuilderPostRequestConfiguration)(UpdatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdatePostResponseable), nil
}
// ToPostRequestInformation updates an existing contact form
func (m *Contact_formUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *Contact_formUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *Contact_formUpdateRequestBuilder) WithUrl(rawUrl string)(*Contact_formUpdateRequestBuilder) {
    return NewContact_formUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
