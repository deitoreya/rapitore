# MailPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailaddress** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewMailPassword

`func NewMailPassword() *MailPassword`

NewMailPassword instantiates a new MailPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailPasswordWithDefaults

`func NewMailPasswordWithDefaults() *MailPassword`

NewMailPasswordWithDefaults instantiates a new MailPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailaddress

`func (o *MailPassword) GetMailaddress() string`

GetMailaddress returns the Mailaddress field if non-nil, zero value otherwise.

### GetMailaddressOk

`func (o *MailPassword) GetMailaddressOk() (*string, bool)`

GetMailaddressOk returns a tuple with the Mailaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailaddress

`func (o *MailPassword) SetMailaddress(v string)`

SetMailaddress sets Mailaddress field to given value.

### HasMailaddress

`func (o *MailPassword) HasMailaddress() bool`

HasMailaddress returns a boolean if a field has been set.

### GetPassword

`func (o *MailPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailPassword) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MailPassword) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


