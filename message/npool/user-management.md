# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/user-management.proto](#npool/user-management.proto)
    - [AddUserRequest](#user.v1.AddUserRequest)
    - [AddUserResponse](#user.v1.AddUserResponse)
    - [BindThirdPartyRequest](#user.v1.BindThirdPartyRequest)
    - [BindThirdPartyResponse](#user.v1.BindThirdPartyResponse)
    - [BindUserEmailRequest](#user.v1.BindUserEmailRequest)
    - [BindUserEmailResponse](#user.v1.BindUserEmailResponse)
    - [BindUserPhoneRequest](#user.v1.BindUserPhoneRequest)
    - [BindUserPhoneResponse](#user.v1.BindUserPhoneResponse)
    - [CertificateKycRequest](#user.v1.CertificateKycRequest)
    - [CertificateKycResponse](#user.v1.CertificateKycResponse)
    - [ChangeUserPasswordRequest](#user.v1.ChangeUserPasswordRequest)
    - [ChangeUserPasswordResponse](#user.v1.ChangeUserPasswordResponse)
    - [DeleteUserRequest](#user.v1.DeleteUserRequest)
    - [DeleteUserResponse](#user.v1.DeleteUserResponse)
    - [ForgetPasswordRequest](#user.v1.ForgetPasswordRequest)
    - [ForgetPasswordResponse](#user.v1.ForgetPasswordResponse)
    - [FrozenUser](#user.v1.FrozenUser)
    - [FrozenUserRequest](#user.v1.FrozenUserRequest)
    - [FrozenUserResponse](#user.v1.FrozenUserResponse)
    - [GetFrozenUsersRequest](#user.v1.GetFrozenUsersRequest)
    - [GetFrozenUsersResponse](#user.v1.GetFrozenUsersResponse)
    - [GetGaQRCodeRequest](#user.v1.GetGaQRCodeRequest)
    - [GetGaQRCodeResponse](#user.v1.GetGaQRCodeResponse)
    - [GetUserProvidersRequest](#user.v1.GetUserProvidersRequest)
    - [GetUserProvidersResponse](#user.v1.GetUserProvidersResponse)
    - [GetUserRequest](#user.v1.GetUserRequest)
    - [GetUserResponse](#user.v1.GetUserResponse)
    - [GetUsersRequest](#user.v1.GetUsersRequest)
    - [GetUsersResponse](#user.v1.GetUsersResponse)
    - [PageInfo](#user.v1.PageInfo)
    - [SignupRequest](#user.v1.SignupRequest)
    - [SignupResponse](#user.v1.SignupResponse)
    - [UnbindThirdPartyRequest](#user.v1.UnbindThirdPartyRequest)
    - [UnbindThirdPartyResponse](#user.v1.UnbindThirdPartyResponse)
    - [UnbindUserEmailRequest](#user.v1.UnbindUserEmailRequest)
    - [UnbindUserEmailResponse](#user.v1.UnbindUserEmailResponse)
    - [UnbindUserPhoneRequest](#user.v1.UnbindUserPhoneRequest)
    - [UnbindUserPhoneResponse](#user.v1.UnbindUserPhoneResponse)
    - [UnfrozenUserRequest](#user.v1.UnfrozenUserRequest)
    - [UnfrozenUserResponse](#user.v1.UnfrozenUserResponse)
    - [UpdateUserInfoRequest](#user.v1.UpdateUserInfoRequest)
    - [UpdateUserInfoResponse](#user.v1.UpdateUserInfoResponse)
    - [UserBasicInfo](#user.v1.UserBasicInfo)
    - [UserProvider](#user.v1.UserProvider)
    - [VerifyGaRequest](#user.v1.VerifyGaRequest)
    - [VerifyGaResponse](#user.v1.VerifyGaResponse)
    - [VersionResponse](#user.v1.VersionResponse)
  
    - [User](#user.v1.User)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/user-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/user-management.proto



<a name="user.v1.AddUserRequest"></a>

### AddUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.AddUserResponse"></a>

### AddUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.BindThirdPartyRequest"></a>

### BindThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| ProviderId | [string](#string) |  | third party(provider)&#39;s id |
| ProviderUserId | [string](#string) |  | UserId in third party(provider) |
| UserProviderInfo | [string](#string) |  |  |






<a name="user.v1.BindThirdPartyResponse"></a>

### BindThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserProviderInfo | [UserProvider](#user.v1.UserProvider) |  |  |






<a name="user.v1.BindUserEmailRequest"></a>

### BindUserEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |






<a name="user.v1.BindUserEmailResponse"></a>

### BindUserEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.BindUserPhoneRequest"></a>

### BindUserPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| PhoneNumber | [string](#string) |  |  |






<a name="user.v1.BindUserPhoneResponse"></a>

### BindUserPhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.CertificateKycRequest"></a>

### CertificateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| FrontCardImg | [string](#string) |  |  |
| BackCardImg | [string](#string) |  |  |
| UserCardImg | [string](#string) |  |  |
| CardType | [string](#string) |  |  |
| CardId | [string](#string) |  |  |
| Region | [string](#string) |  |  |






<a name="user.v1.CertificateKycResponse"></a>

### CertificateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.ChangeUserPasswordRequest"></a>

### ChangeUserPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| OldPassword | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="user.v1.ChangeUserPasswordResponse"></a>

### ChangeUserPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| DeleteUserIds | [string](#string) | repeated | an array of UserId who are being deleted. |






<a name="user.v1.DeleteUserResponse"></a>

### DeleteUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.ForgetPasswordRequest"></a>

### ForgetPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PhoneNumber | [string](#string) |  | Optional |
| EmailAddress | [string](#string) |  | Optional |
| Password | [string](#string) |  |  |






<a name="user.v1.ForgetPasswordResponse"></a>

### ForgetPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.FrozenUser"></a>

### FrozenUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| FrozenBy | [string](#string) |  |  |
| FrozenCause | [string](#string) |  |  |
| StartAt | [uint32](#uint32) |  |  |
| EndAt | [uint32](#uint32) |  |  |
| Status | [string](#string) |  |  |
| UnfrozenBy | [string](#string) |  |  |






<a name="user.v1.FrozenUserRequest"></a>

### FrozenUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| FrozenBy | [string](#string) |  |  |
| FrozenCause | [string](#string) |  |  |






<a name="user.v1.FrozenUserResponse"></a>

### FrozenUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| FrozenUserInfo | [FrozenUser](#user.v1.FrozenUser) |  |  |






<a name="user.v1.GetFrozenUsersRequest"></a>

### GetFrozenUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#user.v1.PageInfo) |  |  |






<a name="user.v1.GetFrozenUsersResponse"></a>

### GetFrozenUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| FrozenUsers | [FrozenUser](#user.v1.FrozenUser) | repeated |  |






<a name="user.v1.GetGaQRCodeRequest"></a>

### GetGaQRCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |






<a name="user.v1.GetGaQRCodeResponse"></a>

### GetGaQRCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GaQRCode | [string](#string) |  |  |






<a name="user.v1.GetUserProvidersRequest"></a>

### GetUserProvidersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |






<a name="user.v1.GetUserProvidersResponse"></a>

### GetUserProvidersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserProviders | [UserProvider](#user.v1.UserProvider) | repeated |  |






<a name="user.v1.GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  | UserId is who is queried. |






<a name="user.v1.GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.GetUsersRequest"></a>

### GetUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#user.v1.PageInfo) |  |  |






<a name="user.v1.GetUsersResponse"></a>

### GetUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfos | [UserBasicInfo](#user.v1.UserBasicInfo) | repeated |  |






<a name="user.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [uint32](#uint32) |  |  |
| PageSize | [uint32](#uint32) |  |  |






<a name="user.v1.SignupRequest"></a>

### SignupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  | optional |
| Password | [string](#string) |  |  |
| EmailAddress | [string](#string) |  | optional |
| PhoneNumber | [string](#string) |  | optional |
| AppId | [string](#string) |  |  |






<a name="user.v1.SignupResponse"></a>

### SignupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UnbindThirdPartyRequest"></a>

### UnbindThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| ProviderId | [string](#string) |  |  |






<a name="user.v1.UnbindThirdPartyResponse"></a>

### UnbindThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserProviderInfo | [UserProvider](#user.v1.UserProvider) |  |  |






<a name="user.v1.UnbindUserEmailRequest"></a>

### UnbindUserEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |






<a name="user.v1.UnbindUserEmailResponse"></a>

### UnbindUserEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.UnbindUserPhoneRequest"></a>

### UnbindUserPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |






<a name="user.v1.UnbindUserPhoneResponse"></a>

### UnbindUserPhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.UnfrozenUserRequest"></a>

### UnfrozenUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| UnfrozenBy | [string](#string) |  |  |






<a name="user.v1.UnfrozenUserResponse"></a>

### UnfrozenUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UnFrozenUserInfo | [FrozenUser](#user.v1.FrozenUser) |  |  |






<a name="user.v1.UpdateUserInfoRequest"></a>

### UpdateUserInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UpdateUserInfoResponse"></a>

### UpdateUserInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UserBasicInfo"></a>

### UserBasicInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Age | [uint32](#uint32) |  |  |
| Gender | [string](#string) |  |  |
| Region | [string](#string) |  |  |
| Birthday | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Province | [string](#string) |  |  |
| City | [string](#string) |  |  |
| PhoneNumber | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |
| LoginTimes | [uint32](#uint32) |  |  |
| KycVerify | [bool](#bool) |  |  |
| GaVerify | [bool](#bool) |  |  |
| SignupMethod | [string](#string) |  |  |
| Career | [string](#string) |  |  |
| DisplayName | [string](#string) |  |  |






<a name="user.v1.UserProvider"></a>

### UserProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| ProviderId | [string](#string) |  |  |
| ProviderUserId | [string](#string) |  |  |
| UserProviderInfo | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |






<a name="user.v1.VerifyGaRequest"></a>

### VerifyGaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) |  |  |
| GaCode | [string](#string) |  |  |






<a name="user.v1.VerifyGaResponse"></a>

### VerifyGaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="user.v1.User"></a>

### User
a service of managing users

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#user.v1.VersionResponse) | Method Version |
| SignUp | [SignupRequest](#user.v1.SignupRequest) | [SignupResponse](#user.v1.SignupResponse) | User can choose signup with username, email, phone or only emial or only phone. |
| GetUser | [GetUserRequest](#user.v1.GetUserRequest) | [GetUserResponse](#user.v1.GetUserResponse) | Get a user&#39;s info by his(her) id, this api can be request by user self of admin. |
| GetUsers | [GetUsersRequest](#user.v1.GetUsersRequest) | [GetUsersResponse](#user.v1.GetUsersResponse) | Get all users. |
| UpdateUserInfo | [UpdateUserInfoRequest](#user.v1.UpdateUserInfoRequest) | [UpdateUserInfoResponse](#user.v1.UpdateUserInfoResponse) | Update user&#39;s basic info. |
| BindUserPhone | [BindUserPhoneRequest](#user.v1.BindUserPhoneRequest) | [BindUserPhoneResponse](#user.v1.BindUserPhoneResponse) | Bind user&#39;s phone number. Before bind a phone number to user, it needs to send phone verify code and confirm the verify code user input. Need apis: https://user.npool.top/v1/send/sms, https://user.npool.top/v1/confirm-sms-verify-code |
| BindUserEmail | [BindUserEmailRequest](#user.v1.BindUserEmailRequest) | [BindUserEmailResponse](#user.v1.BindUserEmailResponse) | Bind user&#39;s email address. Before bind a email address to user, it needs to send email verify code and confirm the verify code user input. Need apis: https://user.npool.top/v1/send/email, https://user.npool.top/v1/confirm-email-verify-code |
| UnbindUserPhone | [UnbindUserPhoneRequest](#user.v1.UnbindUserPhoneRequest) | [UnbindUserPhoneResponse](#user.v1.UnbindUserPhoneResponse) | Unbind user&#39;s phone number. Before unbind user&#39;s phone number, it needs to send phone verify code and confirm the verify code user input. Need apis: https://user.npool.top/v1/send/sms, https://user.npool.top/v1/confirm-sms-verify-code |
| UnbindUserEmail | [UnbindUserEmailRequest](#user.v1.UnbindUserEmailRequest) | [UnbindUserEmailResponse](#user.v1.UnbindUserEmailResponse) | Unbind user&#39;s email address. Before unbind user&#39;s email address, it needs to send email verify code and confirm the verify code user input. Need apis: https://user.npool.top/v1/send/email, https://user.npool.top/v1/confirm-email-verify-code |
| BindThirdParty | [BindThirdPartyRequest](#user.v1.BindThirdPartyRequest) | [BindThirdPartyResponse](#user.v1.BindThirdPartyResponse) | Link to a third-party oauth. save the UserId from third-party into mysql. During Linking, user need to authenticate in third-party and use login-oauth api to get access_token and then get user info from third-party(provider). need api: https://login.npool.top/v1/login/oauth |
| UnbindThirdParty | [UnbindThirdPartyRequest](#user.v1.UnbindThirdPartyRequest) | [UnbindThirdPartyResponse](#user.v1.UnbindThirdPartyResponse) | Unlink a third-party oauth. Delete the UserId we saved from mysql. |
| ChangeUserPassword | [ChangeUserPasswordRequest](#user.v1.ChangeUserPasswordRequest) | [ChangeUserPasswordResponse](#user.v1.ChangeUserPasswordResponse) | Change user&#39;s password. Before change users password, system need the user to do an authentication. |
| ForgetPassword | [ForgetPasswordRequest](#user.v1.ForgetPasswordRequest) | [ForgetPasswordResponse](#user.v1.ForgetPasswordResponse) | Forget password. If user forgets his(her) password, he(she) firstly need to authenticate identity and then reset his(her) password. |
| AddUser | [AddUserRequest](#user.v1.AddUserRequest) | [AddUserResponse](#user.v1.AddUserResponse) | Add user. |
| DeleteUser | [DeleteUserRequest](#user.v1.DeleteUserRequest) | [DeleteUserResponse](#user.v1.DeleteUserResponse) | Delete users. This api can only be used by admin. When deleting users, service will not only delete basic user info, but also use other apis to delete connections among other service. |
| FrozenUser | [FrozenUserRequest](#user.v1.FrozenUserRequest) | [FrozenUserResponse](#user.v1.FrozenUserResponse) | Frozen user. |
| UnfrozenUser | [UnfrozenUserRequest](#user.v1.UnfrozenUserRequest) | [UnfrozenUserResponse](#user.v1.UnfrozenUserResponse) | Unfrozen user. |
| GetFrozenUsers | [GetFrozenUsersRequest](#user.v1.GetFrozenUsersRequest) | [GetFrozenUsersResponse](#user.v1.GetFrozenUsersResponse) | Get frozen user list. |
| GetUserProviders | [GetUserProvidersRequest](#user.v1.GetUserProvidersRequest) | [GetUserProvidersResponse](#user.v1.GetUserProvidersResponse) | Get user providers info. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

