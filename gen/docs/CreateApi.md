# CreateApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**usersCreatePost**](CreateApi.md#usersCreatePost) | **POST** /users/create | CreateUser


<a name="usersCreatePost"></a>
# **usersCreatePost**
> Integer usersCreatePost(input)

CreateUser

create a user

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CreateApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    CreateApi apiInstance = new CreateApi(defaultClient);
    ModelsEmployee input = new ModelsEmployee(); // ModelsEmployee | user info
    try {
      Integer result = apiInstance.usersCreatePost(input);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CreateApi#usersCreatePost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | [**ModelsEmployee**](ModelsEmployee.md)| user info |

### Return type

**Integer**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

