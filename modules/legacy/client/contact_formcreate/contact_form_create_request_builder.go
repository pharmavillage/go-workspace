package contact_formcreate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Contact_formCreateRequestBuilder builds and executes requests for operations under \contact_form.create
type Contact_formCreateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Contact_formCreateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Contact_formCreateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContact_formCreateRequestBuilderInternal instantiates a new Contact_formCreateRequestBuilder and sets the default values.
func NewContact_formCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formCreateRequestBuilder) {
    m := &Contact_formCreateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contact_form.create", pathParameters),
    }
    return m
}
// NewContact_formCreateRequestBuilder instantiates a new Contact_formCreateRequestBuilder and sets the default values.
func NewContact_formCreateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Contact_formCreateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContact_formCreateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new contact form
// Deprecated: This method is obsolete. Use PostAsCreatePostResponse instead.
func (m *Contact_formCreateRequestBuilder) Post(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *Contact_formCreateRequestBuilderPostRequestConfiguration)(CreateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCreate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCreateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CreateResponseable), nil
}
// PostAsCreatePostResponse creates a new contact form
func (m *Contact_formCreateRequestBuilder) PostAsCreatePostResponse(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *Contact_formCreateRequestBuilderPostRequestConfiguration)(CreatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCreate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCreatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CreatePostResponseable), nil
}
// ToPostRequestInformation creates a new contact form
func (m *Contact_formCreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *Contact_formCreateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *Contact_formCreateRequestBuilder) WithUrl(rawUrl string)(*Contact_formCreateRequestBuilder) {
    return NewContact_formCreateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
