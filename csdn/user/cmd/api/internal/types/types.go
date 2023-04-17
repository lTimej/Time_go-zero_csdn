// Code generated by goctl. DO NOT EDIT.
package types

type UserFocusFansList struct {
	UserId       string `json:"user_id"`
	Flag         string `json:"flag"`
	UserName     string `json:"user_name"`
	HeadPhoto    string `json:"head_photo"`
	Introduction string `json:"introduction"`
	MutualFocus  bool   `json:"mutual_focus"`
}

type UserPasswordLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserPasswordLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type PhoneLoginRequest struct {
	Phone string `json:"mobile"`
	Code  string `json:"sms_code"`
}

type PhoneLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SendSmsCodeRequest struct {
}

type SendSmsCodeReponse struct {
}

type UserCurrInfoRequest struct {
}

type UserCurrInfoResponse struct {
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	HeadPhoto string `json:"head_photo"`
	Introduce string `json:"introduce"`
	CodeYear  int32  `json:"code_year"`
	Career    string `json:"career"`
	Focus     int32  `json:"focus"`
	Fans      int32  `json:"fans"`
	Visitor   int32  `json:"visitor"`
}

type IsFocusUserRequest struct {
	TargetUserId string `json:"target_user_id,optional"`
}

type IsFocusUserResponse struct {
	IsFocusUser bool `json:"isFocusUser"`
}

type FocusUserRequest struct {
	TargetUserId string `json:"target"`
}

type FocusUserResponse struct {
	TargetUserId string `json:"target_id,optional"`
}

type CancelFocusUserRequest struct {
	TargetUserId string `json:"target"`
}

type CancelFocusUserResponse struct {
	Message string `json:"message"`
}

type UserFocusListRequest struct {
	Page    int64 `json:"page,optional"`
	PageNum int64 `json:"page_num,optional"`
}

type UserFocusListResponse struct {
	Page     int64                `json:"page"`
	PageNum  int64                `json:"page_num"`
	TotalNum int64                `json:"total_num"`
	Focus    []*UserFocusFansList `json:"focus"`
}

type UserFansListRequest struct {
	Page    int64 `json:"page,optional"`
	PageNum int64 `json:"page_num,optional"`
}

type UserFansListResponse struct {
	Page     int64                `json:"page"`
	PageNum  int64                `json:"page_num"`
	TotalNum int64                `json:"total_num"`
	Fans     []*UserFocusFansList `json:"fans"`
}

type UserInfoEditRequest struct {
	HeadPhoto string `json:"head_photo,optional"`
	OldPwd    string `json:"old_pwd,optional"`
	NewPwd    string `json:"new_pwd,optional"`
	UserName  string `json:"user_name,optional"`
	Gender    string `json:"gender,optional"`
	Introduce string `json:"introduce,optional"`
	Tag       string `json:"tag,optional"`
	AuthName  string `json:"auth_name,optional"`
	Birthday  string `json:"birthday,optional"`
	Areas     string `json:"areas,optional"`
}

type UserInfoEditResponse struct {
	HeadPhoto string `json:"head_photo,optional"`
	Pwd       string `json:"pwd,optional"`
	UserName  string `json:"user_name,optional"`
	Gender    int64  `json:"gender,optional"`
	Introduce string `json:"introduce,optional"`
	Tag       string `json:"tag,optional"`
	AuthName  string `json:"authName,optional"`
	Birthday  string `json:"birthday,optional"`
	Areas     string `json:"areas,optional"`
}

type UserAddressRequest struct {
	Receiver   string `json:"receiver"`
	Mobile     string `json:"mobile"`
	ProvinceId int64  `json:"province_id"`
	CityId     int64  `json:"city_id"`
	DistrictId int64  `json:"district_id"`
	Place      string `json:"place"`
}

type UserAddressResponse struct {
}

type UpdateUserAddressRequest struct {
	AddressId  int64  `json:"address_id"`
	Receiver   string `json:"receiver"`
	Mobile     string `json:"mobile"`
	ProvinceId int64  `json:"province_id"`
	CityId     int64  `json:"city_id"`
	DistrictId int64  `json:"district_id"`
	Place      string `json:"place"`
}

type UpdateUserAddressResponse struct {
}

type UserAddress struct {
	AddressId  int64  `json:"address_id"`
	Receiver   string `json:"receiver"`
	Mobile     string `json:"mobile"`
	ProvinceId int64  `json:"province_id"`
	CityId     int64  `json:"city_id"`
	DistrictId int64  `json:"district_id"`
	Province   string `json:"province"`
	City       string `json:"city"`
	District   string `json:"district"`
	Place      string `json:"place"`
}

type GetUserAddressRequest struct {
}

type GetUserAddressResponse struct {
	UserAddress []*UserAddress `json:"user_address"`
}
