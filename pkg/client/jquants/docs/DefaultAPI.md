# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListedInfoGet**](DefaultAPI.md#ListedInfoGet) | **Get** /listed/info | 
[**TokenAuthRefreshPost**](DefaultAPI.md#TokenAuthRefreshPost) | **Post** /token/auth_refresh | 
[**TokenAuthUserPost**](DefaultAPI.md#TokenAuthUserPost) | **Post** /token/auth_user | 



## ListedInfoGet

> []ListedInfo ListedInfoGet(ctx).Code(code).Date(date).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string |  (optional)
	date := "date_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListedInfoGet(context.Background()).Code(code).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListedInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListedInfoGet`: []ListedInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListedInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListedInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **date** | **string** |  | 

### Return type

[**[]ListedInfo**](ListedInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenAuthRefreshPost

> IDToken TokenAuthRefreshPost(ctx).RefreshToken(refreshToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	refreshToken := *openapiclient.NewRefreshToken() // RefreshToken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TokenAuthRefreshPost(context.Background()).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenAuthRefreshPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenAuthRefreshPost`: IDToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TokenAuthRefreshPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenAuthRefreshPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshToken** | [**RefreshToken**](RefreshToken.md) |  | 

### Return type

[**IDToken**](IDToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenAuthUserPost

> RefreshToken TokenAuthUserPost(ctx).MailPassword(mailPassword).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mailPassword := *openapiclient.NewMailPassword() // MailPassword |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TokenAuthUserPost(context.Background()).MailPassword(mailPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenAuthUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenAuthUserPost`: RefreshToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TokenAuthUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenAuthUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mailPassword** | [**MailPassword**](MailPassword.md) |  | 

### Return type

[**RefreshToken**](RefreshToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

