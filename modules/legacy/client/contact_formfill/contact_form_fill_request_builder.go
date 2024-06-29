package contact_formfill

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Contact_formFillRequestBuilder builds and executes requests for operations under \contact_form.fill
type Contact_formFillRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Contact_formFillRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Contact_formFillRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContact_formFillRequestBuilderInternal instantiates a new Contact_formFillRequestBuilder and sets the default values.
func NewContact_formFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formFillRequestBuilder) {
    m := &Contact_formFillRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contact_form.fill", pathParameters),
    }
    return m
}
// NewContact_formFillRequestBuilder instantiates a new Contact_formFillRequestBuilder and sets the default values.
func NewContact_formFillRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formFillRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContact_formFillRequestBuilderInternal(urlParams, requestAdapter)
}
// Post fills a form, creating a channel and posting the message to it.
// Deprecated: This method is obsolete. Use PostAsFillPostResponse instead.
func (m *Contact_formFillRequestBuilder) Post(ctx context.Context, body FillPostRequestBodyable, requestConfiguration *Contact_formFillRequestBuilderPostRequestConfiguration)(FillResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateFill400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFillResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FillResponseable), nil
}
// PostAsFillPostResponse fills a form, creating a channel and posting the message to it.
func (m *Contact_formFillRequestBuilder) PostAsFillPostResponse(ctx context.Context, body FillPostRequestBodyable, requestConfiguration *Contact_formFillRequestBuilderPostRequestConfiguration)(FillPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateFill400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFillPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FillPostResponseable), nil
}
// ToPostRequestInformation fills a form, creating a channel and posting the message to it.
func (m *Contact_formFillRequestBuilder) ToPostRequestInformation(ctx context.Context, body FillPostRequestBodyable, requestConfiguration *Contact_formFillRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *Contact_formFillRequestBuilder) WithUrl(rawUrl string)(*Contact_formFillRequestBuilder) {
    return NewContact_formFillRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
