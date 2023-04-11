// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Error5xx) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error5xx) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("code")
		e.Int(s.Code)
	}
	{

		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		if s.RequestID.Set {
			e.FieldStart("request_id")
			s.RequestID.Encode(e)
		}
	}
	for k, elem := range s.AdditionalProps {
		e.FieldStart(k)

		elem.Encode(e)
	}
}

var jsonFieldsNameOfError5xx = [3]string{
	0: "code",
	1: "message",
	2: "request_id",
}

// Decode decodes Error5xx from json.
func (s *Error5xx) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error5xx to nil")
	}
	var requiredBitSet [1]uint8
	s.AdditionalProps = map[string]Error5xxAdditionalItem{}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "request_id":
			if err := func() error {
				s.RequestID.Reset()
				if err := s.RequestID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"request_id\"")
			}
		default:
			var elem Error5xxAdditionalItem
			if err := func() error {
				if err := elem.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrapf(err, "decode field %q", k)
			}
			s.AdditionalProps[string(k)] = elem
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error5xx")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError5xx) {
					name = jsonFieldsNameOfError5xx[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error5xx) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error5xx) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Error5xxAdditional) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s Error5xxAdditional) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		elem.Encode(e)
	}
}

// Decode decodes Error5xxAdditional from json.
func (s *Error5xxAdditional) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error5xxAdditional to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem Error5xxAdditionalItem
		if err := func() error {
			if err := elem.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error5xxAdditional")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Error5xxAdditional) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error5xxAdditional) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error5xxAdditionalItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error5xxAdditionalItem) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfError5xxAdditionalItem = [0]string{}

// Decode decodes Error5xxAdditionalItem from json.
func (s *Error5xxAdditionalItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error5xxAdditionalItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error5xxAdditionalItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error5xxAdditionalItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error5xxAdditionalItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ErrorGeneric) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrorGeneric) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("code")
		e.Int(s.Code)
	}
	{

		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		if s.RequestID.Set {
			e.FieldStart("request_id")
			s.RequestID.Encode(e)
		}
	}
	for k, elem := range s.AdditionalProps {
		e.FieldStart(k)

		elem.Encode(e)
	}
}

var jsonFieldsNameOfErrorGeneric = [3]string{
	0: "code",
	1: "message",
	2: "request_id",
}

// Decode decodes ErrorGeneric from json.
func (s *ErrorGeneric) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorGeneric to nil")
	}
	var requiredBitSet [1]uint8
	s.AdditionalProps = map[string]ErrorGenericAdditionalItem{}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "request_id":
			if err := func() error {
				s.RequestID.Reset()
				if err := s.RequestID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"request_id\"")
			}
		default:
			var elem ErrorGenericAdditionalItem
			if err := func() error {
				if err := elem.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrapf(err, "decode field %q", k)
			}
			s.AdditionalProps[string(k)] = elem
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorGeneric")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfErrorGeneric) {
					name = jsonFieldsNameOfErrorGeneric[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrorGeneric) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorGeneric) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ErrorGenericAdditional) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s ErrorGenericAdditional) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		elem.Encode(e)
	}
}

// Decode decodes ErrorGenericAdditional from json.
func (s *ErrorGenericAdditional) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorGenericAdditional to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem ErrorGenericAdditionalItem
		if err := func() error {
			if err := elem.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorGenericAdditional")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ErrorGenericAdditional) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorGenericAdditional) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ErrorGenericAdditionalItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrorGenericAdditionalItem) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfErrorGenericAdditionalItem = [0]string{}

// Decode decodes ErrorGenericAdditionalItem from json.
func (s *ErrorGenericAdditionalItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorGenericAdditionalItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorGenericAdditionalItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrorGenericAdditionalItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorGenericAdditionalItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetUserBadRequest as json.
func (s *GetUserBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorGeneric)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetUserBadRequest from json.
func (s *GetUserBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUserBadRequest to nil")
	}
	var unwrapped ErrorGeneric
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetUserBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUserBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUserBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetUserNotFound as json.
func (s *GetUserNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorGeneric)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetUserNotFound from json.
func (s *GetUserNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUserNotFound to nil")
	}
	var unwrapped ErrorGeneric
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetUserNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUserNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUserNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetUserUnauthorized as json.
func (s *GetUserUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorGeneric)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetUserUnauthorized from json.
func (s *GetUserUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUserUnauthorized to nil")
	}
	var unwrapped ErrorGeneric
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetUserUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUserUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUserUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginInput) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("password")
		e.Str(s.Password)
	}
}

var jsonFieldsNameOfLoginInput = [2]string{
	0: "id",
	1: "password",
}

// Decode decodes LoginInput from json.
func (s *LoginInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLoginInput) {
					name = jsonFieldsNameOfLoginInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginResponse) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("token")
		e.Str(s.Token)
	}
}

var jsonFieldsNameOfLoginResponse = [1]string{
	0: "token",
}

// Decode decodes LoginResponse from json.
func (s *LoginResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "token":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Token = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLoginResponse) {
					name = jsonFieldsNameOfLoginResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes LoginUserBadRequest as json.
func (s *LoginUserBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorGeneric)(s)

	unwrapped.Encode(e)
}

// Decode decodes LoginUserBadRequest from json.
func (s *LoginUserBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginUserBadRequest to nil")
	}
	var unwrapped ErrorGeneric
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = LoginUserBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginUserBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginUserBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes LoginUserNotFound as json.
func (s *LoginUserNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorGeneric)(s)

	unwrapped.Encode(e)
}

// Decode decodes LoginUserNotFound from json.
func (s *LoginUserNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginUserNotFound to nil")
	}
	var unwrapped ErrorGeneric
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = LoginUserNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginUserNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginUserNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int32 as json.
func (o OptInt32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int32(int32(o.Value))
}

// Decode decodes int32 from json.
func (o *OptInt32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt32 to nil")
	}
	o.Set = true
	v, err := d.Int32()
	if err != nil {
		return err
	}
	o.Value = int32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("first_name")
		e.Str(s.FirstName)
	}
	{

		e.FieldStart("second_name")
		e.Str(s.SecondName)
	}
	{

		e.FieldStart("age")
		e.Int32(s.Age)
	}
	{

		e.FieldStart("biography")
		e.Str(s.Biography)
	}
	{

		e.FieldStart("city")
		e.Str(s.City)
	}
}

var jsonFieldsNameOfUser = [6]string{
	0: "id",
	1: "first_name",
	2: "second_name",
	3: "age",
	4: "biography",
	5: "city",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "first_name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.FirstName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"first_name\"")
			}
		case "second_name":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.SecondName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"second_name\"")
			}
		case "age":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int32()
				s.Age = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "biography":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Biography = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "city":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Str()
				s.City = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"city\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00111111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUser) {
					name = jsonFieldsNameOfUser[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserRegister) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserRegister) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("first_name")
		e.Str(s.FirstName)
	}
	{

		e.FieldStart("second_name")
		e.Str(s.SecondName)
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.City.Set {
			e.FieldStart("city")
			s.City.Encode(e)
		}
	}
	{

		e.FieldStart("password")
		e.Str(s.Password)
	}
}

var jsonFieldsNameOfUserRegister = [6]string{
	0: "first_name",
	1: "second_name",
	2: "age",
	3: "biography",
	4: "city",
	5: "password",
}

// Decode decodes UserRegister from json.
func (s *UserRegister) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRegister to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "first_name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.FirstName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"first_name\"")
			}
		case "second_name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.SecondName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"second_name\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "city":
			if err := func() error {
				s.City.Reset()
				if err := s.City.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"city\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserRegister")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00100011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserRegister) {
					name = jsonFieldsNameOfUserRegister[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserRegister) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserRegister) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserRegisterResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserRegisterResponse) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("user_id")
		e.Int64(s.UserID)
	}
}

var jsonFieldsNameOfUserRegisterResponse = [1]string{
	0: "user_id",
}

// Decode decodes UserRegisterResponse from json.
func (s *UserRegisterResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRegisterResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "user_id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.UserID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserRegisterResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserRegisterResponse) {
					name = jsonFieldsNameOfUserRegisterResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserRegisterResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserRegisterResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Users as json.
func (s Users) Encode(e *jx.Encoder) {
	unwrapped := []User(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Users from json.
func (s *Users) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Users to nil")
	}
	var unwrapped []User
	if err := func() error {
		unwrapped = make([]User, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem User
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Users(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Users) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Users) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
