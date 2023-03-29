// Code generated by ogen, DO NOT EDIT.

package generated

type CookieAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *CookieAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *CookieAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type Error5xx struct {
	Code            int       `json:"code"`
	Message         string    `json:"message"`
	RequestID       OptString `json:"request_id"`
	AdditionalProps Error5xxAdditional
}

// GetCode returns the value of Code.
func (s *Error5xx) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error5xx) GetMessage() string {
	return s.Message
}

// GetRequestID returns the value of RequestID.
func (s *Error5xx) GetRequestID() OptString {
	return s.RequestID
}

// GetAdditionalProps returns the value of AdditionalProps.
func (s *Error5xx) GetAdditionalProps() Error5xxAdditional {
	return s.AdditionalProps
}

// SetCode sets the value of Code.
func (s *Error5xx) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error5xx) SetMessage(val string) {
	s.Message = val
}

// SetRequestID sets the value of RequestID.
func (s *Error5xx) SetRequestID(val OptString) {
	s.RequestID = val
}

// SetAdditionalProps sets the value of AdditionalProps.
func (s *Error5xx) SetAdditionalProps(val Error5xxAdditional) {
	s.AdditionalProps = val
}

type Error5xxAdditional map[string]Error5xxAdditionalItem

func (s *Error5xxAdditional) init() Error5xxAdditional {
	m := *s
	if m == nil {
		m = map[string]Error5xxAdditionalItem{}
		*s = m
	}
	return m
}

type Error5xxAdditionalItem struct{}

// Error5xxHeaders wraps Error5xx with response headers.
type Error5xxHeaders struct {
	Retryafter OptInt
	Response   Error5xx
}

// GetRetryafter returns the value of Retryafter.
func (s *Error5xxHeaders) GetRetryafter() OptInt {
	return s.Retryafter
}

// GetResponse returns the value of Response.
func (s *Error5xxHeaders) GetResponse() Error5xx {
	return s.Response
}

// SetRetryafter sets the value of Retryafter.
func (s *Error5xxHeaders) SetRetryafter(val OptInt) {
	s.Retryafter = val
}

// SetResponse sets the value of Response.
func (s *Error5xxHeaders) SetResponse(val Error5xx) {
	s.Response = val
}

func (*Error5xxHeaders) loginUserRes() {}

type ErrorGeneric struct {
	Code            int       `json:"code"`
	Message         string    `json:"message"`
	RequestID       OptString `json:"request_id"`
	AdditionalProps ErrorGenericAdditional
}

// GetCode returns the value of Code.
func (s *ErrorGeneric) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorGeneric) GetMessage() string {
	return s.Message
}

// GetRequestID returns the value of RequestID.
func (s *ErrorGeneric) GetRequestID() OptString {
	return s.RequestID
}

// GetAdditionalProps returns the value of AdditionalProps.
func (s *ErrorGeneric) GetAdditionalProps() ErrorGenericAdditional {
	return s.AdditionalProps
}

// SetCode sets the value of Code.
func (s *ErrorGeneric) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorGeneric) SetMessage(val string) {
	s.Message = val
}

// SetRequestID sets the value of RequestID.
func (s *ErrorGeneric) SetRequestID(val OptString) {
	s.RequestID = val
}

// SetAdditionalProps sets the value of AdditionalProps.
func (s *ErrorGeneric) SetAdditionalProps(val ErrorGenericAdditional) {
	s.AdditionalProps = val
}

func (*ErrorGeneric) registerUserRes()  {}
func (*ErrorGeneric) userSearchGetRes() {}

type ErrorGenericAdditional map[string]ErrorGenericAdditionalItem

func (s *ErrorGenericAdditional) init() ErrorGenericAdditional {
	m := *s
	if m == nil {
		m = map[string]ErrorGenericAdditionalItem{}
		*s = m
	}
	return m
}

type ErrorGenericAdditionalItem struct{}

type GetUserApplicationJSONBadRequest ErrorGeneric

func (*GetUserApplicationJSONBadRequest) getUserRes() {}

type GetUserApplicationJSONInternalServerError Error5xxHeaders

func (*GetUserApplicationJSONInternalServerError) getUserRes() {}

type GetUserApplicationJSONNotFound ErrorGeneric

func (*GetUserApplicationJSONNotFound) getUserRes() {}

type GetUserApplicationJSONServiceUnavailable Error5xxHeaders

func (*GetUserApplicationJSONServiceUnavailable) getUserRes() {}

type GetUserApplicationJSONUnauthorized ErrorGeneric

func (*GetUserApplicationJSONUnauthorized) getUserRes() {}

// Ref: #/components/schemas/LoginInput
type LoginInput struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

// GetID returns the value of ID.
func (s *LoginInput) GetID() int64 {
	return s.ID
}

// GetPassword returns the value of Password.
func (s *LoginInput) GetPassword() string {
	return s.Password
}

// SetID sets the value of ID.
func (s *LoginInput) SetID(val int64) {
	s.ID = val
}

// SetPassword sets the value of Password.
func (s *LoginInput) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/LoginResponse
type LoginResponse struct {
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *LoginResponse) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *LoginResponse) SetToken(val string) {
	s.Token = val
}

// LoginResponseHeaders wraps LoginResponse with response headers.
type LoginResponseHeaders struct {
	SetCookie string
	Response  LoginResponse
}

// GetSetCookie returns the value of SetCookie.
func (s *LoginResponseHeaders) GetSetCookie() string {
	return s.SetCookie
}

// GetResponse returns the value of Response.
func (s *LoginResponseHeaders) GetResponse() LoginResponse {
	return s.Response
}

// SetSetCookie sets the value of SetCookie.
func (s *LoginResponseHeaders) SetSetCookie(val string) {
	s.SetCookie = val
}

// SetResponse sets the value of Response.
func (s *LoginResponseHeaders) SetResponse(val LoginResponse) {
	s.Response = val
}

func (*LoginResponseHeaders) loginUserRes() {}

type LoginUserApplicationJSONBadRequest ErrorGeneric

func (*LoginUserApplicationJSONBadRequest) loginUserRes() {}

type LoginUserApplicationJSONNotFound ErrorGeneric

func (*LoginUserApplicationJSONNotFound) loginUserRes() {}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type RegisterUserApplicationJSONInternalServerError Error5xxHeaders

func (*RegisterUserApplicationJSONInternalServerError) registerUserRes() {}

type RegisterUserApplicationJSONServiceUnavailable Error5xxHeaders

func (*RegisterUserApplicationJSONServiceUnavailable) registerUserRes() {}

// Ref: #/components/schemas/User
type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int32  `json:"age"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
}

// GetID returns the value of ID.
func (s *User) GetID() int64 {
	return s.ID
}

// GetFirstName returns the value of FirstName.
func (s *User) GetFirstName() string {
	return s.FirstName
}

// GetSecondName returns the value of SecondName.
func (s *User) GetSecondName() string {
	return s.SecondName
}

// GetAge returns the value of Age.
func (s *User) GetAge() int32 {
	return s.Age
}

// GetBiography returns the value of Biography.
func (s *User) GetBiography() string {
	return s.Biography
}

// GetCity returns the value of City.
func (s *User) GetCity() string {
	return s.City
}

// SetID sets the value of ID.
func (s *User) SetID(val int64) {
	s.ID = val
}

// SetFirstName sets the value of FirstName.
func (s *User) SetFirstName(val string) {
	s.FirstName = val
}

// SetSecondName sets the value of SecondName.
func (s *User) SetSecondName(val string) {
	s.SecondName = val
}

// SetAge sets the value of Age.
func (s *User) SetAge(val int32) {
	s.Age = val
}

// SetBiography sets the value of Biography.
func (s *User) SetBiography(val string) {
	s.Biography = val
}

// SetCity sets the value of City.
func (s *User) SetCity(val string) {
	s.City = val
}

func (*User) getUserRes() {}

// Ref: #/components/schemas/UserRegister
type UserRegister struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Age        OptInt32  `json:"age"`
	Biography  OptString `json:"biography"`
	City       OptString `json:"city"`
	Password   string    `json:"password"`
}

// GetFirstName returns the value of FirstName.
func (s *UserRegister) GetFirstName() string {
	return s.FirstName
}

// GetSecondName returns the value of SecondName.
func (s *UserRegister) GetSecondName() string {
	return s.SecondName
}

// GetAge returns the value of Age.
func (s *UserRegister) GetAge() OptInt32 {
	return s.Age
}

// GetBiography returns the value of Biography.
func (s *UserRegister) GetBiography() OptString {
	return s.Biography
}

// GetCity returns the value of City.
func (s *UserRegister) GetCity() OptString {
	return s.City
}

// GetPassword returns the value of Password.
func (s *UserRegister) GetPassword() string {
	return s.Password
}

// SetFirstName sets the value of FirstName.
func (s *UserRegister) SetFirstName(val string) {
	s.FirstName = val
}

// SetSecondName sets the value of SecondName.
func (s *UserRegister) SetSecondName(val string) {
	s.SecondName = val
}

// SetAge sets the value of Age.
func (s *UserRegister) SetAge(val OptInt32) {
	s.Age = val
}

// SetBiography sets the value of Biography.
func (s *UserRegister) SetBiography(val OptString) {
	s.Biography = val
}

// SetCity sets the value of City.
func (s *UserRegister) SetCity(val OptString) {
	s.City = val
}

// SetPassword sets the value of Password.
func (s *UserRegister) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/UserRegisterResponse
type UserRegisterResponse struct {
	UserID int64 `json:"user_id"`
}

// GetUserID returns the value of UserID.
func (s *UserRegisterResponse) GetUserID() int64 {
	return s.UserID
}

// SetUserID sets the value of UserID.
func (s *UserRegisterResponse) SetUserID(val int64) {
	s.UserID = val
}

func (*UserRegisterResponse) registerUserRes() {}

type UserSearchGetApplicationJSONInternalServerError Error5xxHeaders

func (*UserSearchGetApplicationJSONInternalServerError) userSearchGetRes() {}

type UserSearchGetApplicationJSONServiceUnavailable Error5xxHeaders

func (*UserSearchGetApplicationJSONServiceUnavailable) userSearchGetRes() {}

type Users []User

func (*Users) userSearchGetRes() {}
