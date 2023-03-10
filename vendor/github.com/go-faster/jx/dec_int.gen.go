// Code generated by mkint, DO NOT EDIT.

package jx

import (
	"io"
	"math"
	"strconv"

	"github.com/go-faster/errors"
)

var errOverflow = strconv.ErrRange

const (
	uint8SafeToMultiple10  = uint8(0xff)/10 - 1
	uint16SafeToMultiple10 = uint16(0xffff)/10 - 1
	uint32SafeToMultiple10 = uint32(0xffffffff)/10 - 1
	uint64SafeToMultiple10 = uint64(0xffffffffffffffff)/10 - 1
)

// UInt8 reads uint8.
func (d *Decoder) UInt8() (uint8, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	return d.readUInt8(c)
}

func (d *Decoder) readUInt8(c byte) (uint8, error) {
	ind := floatDigits[c]
	switch ind {
	case 0:
		// Check that next byte is not a digit.
		c, err := d.peek()
		if err == nil {
			switch floatDigits[c] {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "digit after leading zero")
			case dotInNumber, expInNumber, plusInNumber, minusInNumber:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "unexpected floating point character")
			case invalidCharForNumber:
				return 0, badToken(c, d.offset())
			}
		}
		return 0, nil // single zero
	default:
		if ind < 0 {
			return 0, badToken(c, d.offset()-1)
		}
	}
	value := uint8(ind)
	if d.tail-d.head > 3 {
		i := d.head
		// Iteration 0.
		ind2 := floatDigits[d.buf[i]]
		switch ind2 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+0)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+0)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1
			return value, nil
		}
		i++
		// Iteration 1.
		ind3 := floatDigits[d.buf[i]]
		switch ind3 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+1)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+1)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10
			value += uint8(ind2) * 1
			return value, nil
		}
		i++
		// Iteration 2.
		ind4 := floatDigits[d.buf[i]]
		switch ind4 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+2)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+2)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100
			value += uint8(ind2) * 10
			value += uint8(ind3) * 1
			return value, nil
		}
		d.head = i
		value *= 100
		value += uint8(ind2) * 10
		value += uint8(ind3) * 1
	}
	for {
		buf := d.buf[d.head:d.tail]
		for i, c := range buf {
			ind = floatDigits[c]
			switch ind {
			case invalidCharForNumber:
				return 0, badToken(c, d.offset()+i)
			case dotInNumber,
				expInNumber,
				plusInNumber,
				minusInNumber:
				err := badToken(c, d.offset()+i)
				return 0, errors.Wrap(err, "unexpected floating point character")
			case endOfNumber:
				d.head += i
				return value, nil
			}
			if value > uint8SafeToMultiple10 {
				value2 := (value << 3) + (value << 1) + uint8(ind)
				if value2 < value {
					return 0, errOverflow
				}
				value = value2
				continue
			}
			value = (value << 3) + (value << 1) + uint8(ind)
		}
		switch err := d.read(); err {
		case io.EOF:
			return value, nil
		case nil:
			continue
		default:
			return 0, err
		}
	}
}

// Int8 reads int8.
func (d *Decoder) Int8() (int8, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	if c == '-' {
		c, err := d.byte()
		if err != nil {
			return 0, err
		}
		val, err := d.readUInt8(c)
		if err != nil {
			return 0, err
		}
		if val > math.MaxInt8+1 {
			return 0, errOverflow
		}
		return -int8(val), nil
	}
	val, err := d.readUInt8(c)
	if err != nil {
		return 0, err
	}
	if val > math.MaxInt8 {
		return 0, errOverflow
	}
	return int8(val), nil
}

// UInt16 reads uint16.
func (d *Decoder) UInt16() (uint16, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	return d.readUInt16(c)
}

func (d *Decoder) readUInt16(c byte) (uint16, error) {
	ind := floatDigits[c]
	switch ind {
	case 0:
		// Check that next byte is not a digit.
		c, err := d.peek()
		if err == nil {
			switch floatDigits[c] {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "digit after leading zero")
			case dotInNumber, expInNumber, plusInNumber, minusInNumber:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "unexpected floating point character")
			case invalidCharForNumber:
				return 0, badToken(c, d.offset())
			}
		}
		return 0, nil // single zero
	default:
		if ind < 0 {
			return 0, badToken(c, d.offset()-1)
		}
	}
	value := uint16(ind)
	if d.tail-d.head > 5 {
		i := d.head
		// Iteration 0.
		ind2 := floatDigits[d.buf[i]]
		switch ind2 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+0)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+0)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1
			return value, nil
		}
		i++
		// Iteration 1.
		ind3 := floatDigits[d.buf[i]]
		switch ind3 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+1)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+1)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10
			value += uint16(ind2) * 1
			return value, nil
		}
		i++
		// Iteration 2.
		ind4 := floatDigits[d.buf[i]]
		switch ind4 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+2)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+2)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100
			value += uint16(ind2) * 10
			value += uint16(ind3) * 1
			return value, nil
		}
		i++
		// Iteration 3.
		ind5 := floatDigits[d.buf[i]]
		switch ind5 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+3)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+3)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1000
			value += uint16(ind2) * 100
			value += uint16(ind3) * 10
			value += uint16(ind4) * 1
			return value, nil
		}
		i++
		// Iteration 4.
		ind6 := floatDigits[d.buf[i]]
		switch ind6 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+4)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+4)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10000
			value += uint16(ind2) * 1000
			value += uint16(ind3) * 100
			value += uint16(ind4) * 10
			value += uint16(ind5) * 1
			return value, nil
		}
		d.head = i
		value *= 10000
		value += uint16(ind2) * 1000
		value += uint16(ind3) * 100
		value += uint16(ind4) * 10
		value += uint16(ind5) * 1
	}
	for {
		buf := d.buf[d.head:d.tail]
		for i, c := range buf {
			ind = floatDigits[c]
			switch ind {
			case invalidCharForNumber:
				return 0, badToken(c, d.offset()+i)
			case dotInNumber,
				expInNumber,
				plusInNumber,
				minusInNumber:
				err := badToken(c, d.offset()+i)
				return 0, errors.Wrap(err, "unexpected floating point character")
			case endOfNumber:
				d.head += i
				return value, nil
			}
			if value > uint16SafeToMultiple10 {
				value2 := (value << 3) + (value << 1) + uint16(ind)
				if value2 < value {
					return 0, errOverflow
				}
				value = value2
				continue
			}
			value = (value << 3) + (value << 1) + uint16(ind)
		}
		switch err := d.read(); err {
		case io.EOF:
			return value, nil
		case nil:
			continue
		default:
			return 0, err
		}
	}
}

// Int16 reads int16.
func (d *Decoder) Int16() (int16, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	if c == '-' {
		c, err := d.byte()
		if err != nil {
			return 0, err
		}
		val, err := d.readUInt16(c)
		if err != nil {
			return 0, err
		}
		if val > math.MaxInt16+1 {
			return 0, errOverflow
		}
		return -int16(val), nil
	}
	val, err := d.readUInt16(c)
	if err != nil {
		return 0, err
	}
	if val > math.MaxInt16 {
		return 0, errOverflow
	}
	return int16(val), nil
}

// UInt32 reads uint32.
func (d *Decoder) UInt32() (uint32, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	return d.readUInt32(c)
}

func (d *Decoder) readUInt32(c byte) (uint32, error) {
	ind := floatDigits[c]
	switch ind {
	case 0:
		// Check that next byte is not a digit.
		c, err := d.peek()
		if err == nil {
			switch floatDigits[c] {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "digit after leading zero")
			case dotInNumber, expInNumber, plusInNumber, minusInNumber:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "unexpected floating point character")
			case invalidCharForNumber:
				return 0, badToken(c, d.offset())
			}
		}
		return 0, nil // single zero
	default:
		if ind < 0 {
			return 0, badToken(c, d.offset()-1)
		}
	}
	value := uint32(ind)
	if d.tail-d.head > 9 {
		i := d.head
		// Iteration 0.
		ind2 := floatDigits[d.buf[i]]
		switch ind2 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+0)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+0)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1
			return value, nil
		}
		i++
		// Iteration 1.
		ind3 := floatDigits[d.buf[i]]
		switch ind3 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+1)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+1)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10
			value += uint32(ind2) * 1
			return value, nil
		}
		i++
		// Iteration 2.
		ind4 := floatDigits[d.buf[i]]
		switch ind4 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+2)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+2)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100
			value += uint32(ind2) * 10
			value += uint32(ind3) * 1
			return value, nil
		}
		i++
		// Iteration 3.
		ind5 := floatDigits[d.buf[i]]
		switch ind5 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+3)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+3)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1000
			value += uint32(ind2) * 100
			value += uint32(ind3) * 10
			value += uint32(ind4) * 1
			return value, nil
		}
		i++
		// Iteration 4.
		ind6 := floatDigits[d.buf[i]]
		switch ind6 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+4)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+4)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10000
			value += uint32(ind2) * 1000
			value += uint32(ind3) * 100
			value += uint32(ind4) * 10
			value += uint32(ind5) * 1
			return value, nil
		}
		i++
		// Iteration 5.
		ind7 := floatDigits[d.buf[i]]
		switch ind7 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+5)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+5)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100000
			value += uint32(ind2) * 10000
			value += uint32(ind3) * 1000
			value += uint32(ind4) * 100
			value += uint32(ind5) * 10
			value += uint32(ind6) * 1
			return value, nil
		}
		i++
		// Iteration 6.
		ind8 := floatDigits[d.buf[i]]
		switch ind8 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+6)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+6)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1000000
			value += uint32(ind2) * 100000
			value += uint32(ind3) * 10000
			value += uint32(ind4) * 1000
			value += uint32(ind5) * 100
			value += uint32(ind6) * 10
			value += uint32(ind7) * 1
			return value, nil
		}
		i++
		// Iteration 7.
		ind9 := floatDigits[d.buf[i]]
		switch ind9 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+7)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+7)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10000000
			value += uint32(ind2) * 1000000
			value += uint32(ind3) * 100000
			value += uint32(ind4) * 10000
			value += uint32(ind5) * 1000
			value += uint32(ind6) * 100
			value += uint32(ind7) * 10
			value += uint32(ind8) * 1
			return value, nil
		}
		i++
		// Iteration 8.
		ind10 := floatDigits[d.buf[i]]
		switch ind10 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+8)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+8)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100000000
			value += uint32(ind2) * 10000000
			value += uint32(ind3) * 1000000
			value += uint32(ind4) * 100000
			value += uint32(ind5) * 10000
			value += uint32(ind6) * 1000
			value += uint32(ind7) * 100
			value += uint32(ind8) * 10
			value += uint32(ind9) * 1
			return value, nil
		}
		d.head = i
		value *= 100000000
		value += uint32(ind2) * 10000000
		value += uint32(ind3) * 1000000
		value += uint32(ind4) * 100000
		value += uint32(ind5) * 10000
		value += uint32(ind6) * 1000
		value += uint32(ind7) * 100
		value += uint32(ind8) * 10
		value += uint32(ind9) * 1
	}
	for {
		buf := d.buf[d.head:d.tail]
		for i, c := range buf {
			ind = floatDigits[c]
			switch ind {
			case invalidCharForNumber:
				return 0, badToken(c, d.offset()+i)
			case dotInNumber,
				expInNumber,
				plusInNumber,
				minusInNumber:
				err := badToken(c, d.offset()+i)
				return 0, errors.Wrap(err, "unexpected floating point character")
			case endOfNumber:
				d.head += i
				return value, nil
			}
			if value > uint32SafeToMultiple10 {
				value2 := (value << 3) + (value << 1) + uint32(ind)
				if value2 < value {
					return 0, errOverflow
				}
				value = value2
				continue
			}
			value = (value << 3) + (value << 1) + uint32(ind)
		}
		switch err := d.read(); err {
		case io.EOF:
			return value, nil
		case nil:
			continue
		default:
			return 0, err
		}
	}
}

// Int32 reads int32.
func (d *Decoder) Int32() (int32, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	if c == '-' {
		c, err := d.byte()
		if err != nil {
			return 0, err
		}
		val, err := d.readUInt32(c)
		if err != nil {
			return 0, err
		}
		if val > math.MaxInt32+1 {
			return 0, errOverflow
		}
		return -int32(val), nil
	}
	val, err := d.readUInt32(c)
	if err != nil {
		return 0, err
	}
	if val > math.MaxInt32 {
		return 0, errOverflow
	}
	return int32(val), nil
}

// UInt64 reads uint64.
func (d *Decoder) UInt64() (uint64, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	return d.readUInt64(c)
}

func (d *Decoder) readUInt64(c byte) (uint64, error) {
	ind := floatDigits[c]
	switch ind {
	case 0:
		// Check that next byte is not a digit.
		c, err := d.peek()
		if err == nil {
			switch floatDigits[c] {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "digit after leading zero")
			case dotInNumber, expInNumber, plusInNumber, minusInNumber:
				err := badToken(c, d.offset())
				return 0, errors.Wrap(err, "unexpected floating point character")
			case invalidCharForNumber:
				return 0, badToken(c, d.offset())
			}
		}
		return 0, nil // single zero
	default:
		if ind < 0 {
			return 0, badToken(c, d.offset()-1)
		}
	}
	value := uint64(ind)
	if d.tail-d.head > 9 {
		i := d.head
		// Iteration 0.
		ind2 := floatDigits[d.buf[i]]
		switch ind2 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+0)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+0)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1
			return value, nil
		}
		i++
		// Iteration 1.
		ind3 := floatDigits[d.buf[i]]
		switch ind3 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+1)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+1)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10
			value += uint64(ind2) * 1
			return value, nil
		}
		i++
		// Iteration 2.
		ind4 := floatDigits[d.buf[i]]
		switch ind4 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+2)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+2)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100
			value += uint64(ind2) * 10
			value += uint64(ind3) * 1
			return value, nil
		}
		i++
		// Iteration 3.
		ind5 := floatDigits[d.buf[i]]
		switch ind5 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+3)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+3)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1000
			value += uint64(ind2) * 100
			value += uint64(ind3) * 10
			value += uint64(ind4) * 1
			return value, nil
		}
		i++
		// Iteration 4.
		ind6 := floatDigits[d.buf[i]]
		switch ind6 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+4)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+4)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10000
			value += uint64(ind2) * 1000
			value += uint64(ind3) * 100
			value += uint64(ind4) * 10
			value += uint64(ind5) * 1
			return value, nil
		}
		i++
		// Iteration 5.
		ind7 := floatDigits[d.buf[i]]
		switch ind7 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+5)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+5)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100000
			value += uint64(ind2) * 10000
			value += uint64(ind3) * 1000
			value += uint64(ind4) * 100
			value += uint64(ind5) * 10
			value += uint64(ind6) * 1
			return value, nil
		}
		i++
		// Iteration 6.
		ind8 := floatDigits[d.buf[i]]
		switch ind8 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+6)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+6)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 1000000
			value += uint64(ind2) * 100000
			value += uint64(ind3) * 10000
			value += uint64(ind4) * 1000
			value += uint64(ind5) * 100
			value += uint64(ind6) * 10
			value += uint64(ind7) * 1
			return value, nil
		}
		i++
		// Iteration 7.
		ind9 := floatDigits[d.buf[i]]
		switch ind9 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+7)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+7)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 10000000
			value += uint64(ind2) * 1000000
			value += uint64(ind3) * 100000
			value += uint64(ind4) * 10000
			value += uint64(ind5) * 1000
			value += uint64(ind6) * 100
			value += uint64(ind7) * 10
			value += uint64(ind8) * 1
			return value, nil
		}
		i++
		// Iteration 8.
		ind10 := floatDigits[d.buf[i]]
		switch ind10 {
		case invalidCharForNumber:
			return 0, badToken(d.buf[i], d.offset()+8)
		case dotInNumber,
			expInNumber,
			plusInNumber,
			minusInNumber:
			err := badToken(d.buf[i], d.offset()+8)
			return 0, errors.Wrap(err, "unexpected floating point character")
		case endOfNumber:
			d.head = i
			value *= 100000000
			value += uint64(ind2) * 10000000
			value += uint64(ind3) * 1000000
			value += uint64(ind4) * 100000
			value += uint64(ind5) * 10000
			value += uint64(ind6) * 1000
			value += uint64(ind7) * 100
			value += uint64(ind8) * 10
			value += uint64(ind9) * 1
			return value, nil
		}
		d.head = i
		value *= 100000000
		value += uint64(ind2) * 10000000
		value += uint64(ind3) * 1000000
		value += uint64(ind4) * 100000
		value += uint64(ind5) * 10000
		value += uint64(ind6) * 1000
		value += uint64(ind7) * 100
		value += uint64(ind8) * 10
		value += uint64(ind9) * 1
	}
	for {
		buf := d.buf[d.head:d.tail]
		for i, c := range buf {
			ind = floatDigits[c]
			switch ind {
			case invalidCharForNumber:
				return 0, badToken(c, d.offset()+i)
			case dotInNumber,
				expInNumber,
				plusInNumber,
				minusInNumber:
				err := badToken(c, d.offset()+i)
				return 0, errors.Wrap(err, "unexpected floating point character")
			case endOfNumber:
				d.head += i
				return value, nil
			}
			if value > uint64SafeToMultiple10 {
				value2 := (value << 3) + (value << 1) + uint64(ind)
				if value2 < value {
					return 0, errOverflow
				}
				value = value2
				continue
			}
			value = (value << 3) + (value << 1) + uint64(ind)
		}
		switch err := d.read(); err {
		case io.EOF:
			return value, nil
		case nil:
			continue
		default:
			return 0, err
		}
	}
}

// Int64 reads int64.
func (d *Decoder) Int64() (int64, error) {
	c, err := d.more()
	if err != nil {
		return 0, err
	}
	if c == '-' {
		c, err := d.byte()
		if err != nil {
			return 0, err
		}
		val, err := d.readUInt64(c)
		if err != nil {
			return 0, err
		}
		if val > math.MaxInt64+1 {
			return 0, errOverflow
		}
		return -int64(val), nil
	}
	val, err := d.readUInt64(c)
	if err != nil {
		return 0, err
	}
	if val > math.MaxInt64 {
		return 0, errOverflow
	}
	return int64(val), nil
}
