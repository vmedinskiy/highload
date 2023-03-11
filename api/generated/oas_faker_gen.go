// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"fmt"
)

// SetFake set fake values.
func (s *Error5xx) SetFake() {
	{
		{
			s.Code.SetFake()
		}
	}
	{
		{
			s.Message = "string"
		}
	}
	{
		{
			s.RequestID.SetFake()
		}
	}
	{
		{
			s.AdditionalProps.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Error5xxAdditional) SetFake() {
	var (
		elem Error5xxAdditionalItem
		m    map[string]Error5xxAdditionalItem = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *Error5xxAdditionalItem) SetFake() {
}

// SetFake set fake values.
func (s *ErrorGeneric) SetFake() {
	{
		{
			s.Code.SetFake()
		}
	}
	{
		{
			s.Message = "string"
		}
	}
	{
		{
			s.RequestID.SetFake()
		}
	}
	{
		{
			s.AdditionalProps.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ErrorGenericAdditional) SetFake() {
	var (
		elem ErrorGenericAdditionalItem
		m    map[string]ErrorGenericAdditionalItem = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *ErrorGenericAdditionalItem) SetFake() {
}

// SetFake set fake values.
func (s *GetUserApplicationJSONBadRequest) SetFake() {
	var unwrapped ErrorGeneric
	{
		unwrapped.SetFake()
	}
	*s = GetUserApplicationJSONBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *GetUserApplicationJSONNotFound) SetFake() {
	var unwrapped ErrorGeneric
	{
		unwrapped.SetFake()
	}
	*s = GetUserApplicationJSONNotFound(unwrapped)
}

// SetFake set fake values.
func (s *GetUserApplicationJSONUnauthorized) SetFake() {
	var unwrapped ErrorGeneric
	{
		unwrapped.SetFake()
	}
	*s = GetUserApplicationJSONUnauthorized(unwrapped)
}

// SetFake set fake values.
func (s *LoginInput) SetFake() {
	{
		{
			s.ID = int64(0)
		}
	}
	{
		{
			s.Password = "string"
		}
	}
}

// SetFake set fake values.
func (s *LoginResponse) SetFake() {
	{
		{
			s.Token = "string"
		}
	}
}

// SetFake set fake values.
func (s *LoginUserApplicationJSONBadRequest) SetFake() {
	var unwrapped ErrorGeneric
	{
		unwrapped.SetFake()
	}
	*s = LoginUserApplicationJSONBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *LoginUserApplicationJSONNotFound) SetFake() {
	var unwrapped ErrorGeneric
	{
		unwrapped.SetFake()
	}
	*s = LoginUserApplicationJSONNotFound(unwrapped)
}

// SetFake set fake values.
func (s *OptInt) SetFake() {
	var elem int
	{
		elem = int(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt64) SetFake() {
	var elem int64
	{
		elem = int64(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptString) SetFake() {
	var elem string
	{
		elem = "string"
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptUserRegister) SetFake() {
	var elem UserRegister
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *User) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.FirstName.SetFake()
		}
	}
	{
		{
			s.SecondName.SetFake()
		}
	}
	{
		{
			s.Age.SetFake()
		}
	}
	{
		{
			s.Biography.SetFake()
		}
	}
	{
		{
			s.City.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *UserRegister) SetFake() {
	{
		{
			s.FirstName = "string"
		}
	}
	{
		{
			s.SecondName = "string"
		}
	}
	{
		{
			s.Age.SetFake()
		}
	}
	{
		{
			s.Biography.SetFake()
		}
	}
	{
		{
			s.City.SetFake()
		}
	}
	{
		{
			s.Password = "string"
		}
	}
}

// SetFake set fake values.
func (s *UserRegisterResponse) SetFake() {
	{
		{
			s.UserID.SetFake()
		}
	}
}
