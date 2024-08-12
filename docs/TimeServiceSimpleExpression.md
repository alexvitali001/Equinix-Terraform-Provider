# TimeServiceSimpleExpression

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** | Possible field names to use on filters:  * &#x60;/project/projectId&#x60; - project id (mandatory)  * &#x60;/name&#x60; - Precision Time Service name  * &#x60;/uuid&#x60; - Precision Time Service uuid  * &#x60;/type&#x60; - Precision Time Service protocol  * &#x60;/state&#x60; - Precision Time Service status  * &#x60;/account/accountNumber&#x60; - Precision Time Service account number  * &#x60;/package/code&#x60; - Precision Time Service package  * &#x60;/_*&#x60; - all-category search  | [optional] [default to null]
**Operator** | **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] [default to null]
**Values** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

