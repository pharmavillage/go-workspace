package searchquery

import (
    "context"
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SearchQueryRequestBuilder builds and executes requests for operations under \search.query
type SearchQueryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SearchQueryRequestBuilderGetQueryParameters search anywhere based on a search query
type SearchQueryRequestBuilderGetQueryParameters struct {
    // Search only on one particular channel. If not provided, includes all channels.
    Channel *int32 `uriparametername:"channel"`
    // Number of results to retrieve (per type). Defaults to 3.
    Limit *int32 `uriparametername:"limit"`
    // The search query
    Query *string `uriparametername:"query"`
    // Limit the scope of the search based on the result type. If not provided, all result types are included.
    // Deprecated: This property is deprecated, use scopeAsSearchTypes instead
    Scope *string `uriparametername:"scope"`
    // Limit the scope of the search based on the result type. If not provided, all result types are included.
    ScopeAsSearchTypes *i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchTypes `uriparametername:"scope"`
}
// SearchQueryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchQueryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SearchQueryRequestBuilderGetQueryParameters
}
// NewSearchQueryRequestBuilderInternal instantiates a new SearchQueryRequestBuilder and sets the default values.
func NewSearchQueryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchQueryRequestBuilder) {
    m := &SearchQueryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search.query?query={query}{&channel*,limit*,scope*}", pathParameters),
    }
    return m
}
// NewSearchQueryRequestBuilder instantiates a new SearchQueryRequestBuilder and sets the default values.
func NewSearchQueryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchQueryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchQueryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get search anywhere based on a search query
// Deprecated: This method is obsolete. Use GetAsQueryGetResponse instead.
func (m *SearchQueryRequestBuilder) Get(ctx context.Context, requestConfiguration *SearchQueryRequestBuilderGetRequestConfiguration)(QueryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateQuery400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateQueryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(QueryResponseable), nil
}
// GetAsQueryGetResponse search anywhere based on a search query
func (m *SearchQueryRequestBuilder) GetAsQueryGetResponse(ctx context.Context, requestConfiguration *SearchQueryRequestBuilderGetRequestConfiguration)(QueryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateQuery400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateQueryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(QueryGetResponseable), nil
}
// ToGetRequestInformation search anywhere based on a search query
func (m *SearchQueryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SearchQueryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SearchQueryRequestBuilder) WithUrl(rawUrl string)(*SearchQueryRequestBuilder) {
    return NewSearchQueryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
