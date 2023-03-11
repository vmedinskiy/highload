package uri

import (
	"fmt"
	"net/url"
	"strings"
)

type PathStyle string

const (
	PathStyleSimple PathStyle = "simple"
	PathStyleLabel  PathStyle = "label"
	PathStyleMatrix PathStyle = "matrix"
)

func (s PathStyle) String() string { return string(s) }

type PathEncoder struct {
	param   string    // immutable
	style   PathStyle // immutable
	explode bool      // immutable
	*receiver
}

type PathEncoderConfig struct {
	Param   string
	Style   PathStyle
	Explode bool
}

func NewPathEncoder(cfg PathEncoderConfig) *PathEncoder {
	return &PathEncoder{
		receiver: newReceiver(),
		param:    cfg.Param,
		style:    cfg.Style,
		explode:  cfg.Explode,
	}
}

func (e *PathEncoder) Result() (r string) {
	switch e.typ {
	case typeNotSet:
		panic("encoder was not called, no value")
	case typeValue:
		e.val = url.PathEscape(e.val)
		return e.value()
	case typeArray:
		for i := range e.items {
			e.items[i] = url.PathEscape(e.items[i])
		}
		return e.array()
	case typeObject:
		for i, f := range e.fields {
			e.fields[i] = Field{
				Name:  url.PathEscape(f.Name),
				Value: url.PathEscape(f.Value),
			}
		}
		return e.object()
	default:
		panic(fmt.Sprintf("unexpected value type: %T", e.typ))
	}
}

func (e *PathEncoder) value() string {
	switch e.style {
	case PathStyleSimple:
		return e.val
	case PathStyleLabel:
		return "." + e.val
	case PathStyleMatrix:
		return ";" + e.param + "=" + e.val
	default:
		panic("unreachable")
	}
}

func (e *PathEncoder) array() string {
	switch e.style {
	case PathStyleSimple:
		var result []rune
		ll := len(e.items)
		for i := 0; i < ll; i++ {
			result = append(result, []rune(e.items[i])...)
			if i != ll-1 {
				result = append(result, ',')
			}
		}
		return string(result)
	case PathStyleLabel:
		if !e.explode {
			return "." + strings.Join(e.items, ",")
		}
		return "." + strings.Join(e.items, ".")
	case PathStyleMatrix:
		if !e.explode {
			return ";" + e.param + "=" + strings.Join(e.items, ",")
		}
		prefix := ";" + e.param + "="
		return prefix + strings.Join(e.items, prefix)
	default:
		panic("unreachable")
	}
}

func (e *PathEncoder) object() string {
	switch e.style {
	case PathStyleSimple:
		if e.explode {
			const kvSep, fieldSep = '=', ','
			return encodeObject(kvSep, fieldSep, e.fields)
		}

		const kvSep, fieldSep = ',', ','
		return encodeObject(kvSep, fieldSep, e.fields)

	case PathStyleLabel:
		var kvSep, fieldSep byte = ',', ','
		if e.explode {
			kvSep, fieldSep = '=', '.'
		}
		return "." + encodeObject(kvSep, fieldSep, e.fields)

	case PathStyleMatrix:
		if !e.explode {
			const kvSep, fieldSep = ',', ','
			return ";" + e.param + "=" + encodeObject(kvSep, fieldSep, e.fields)
		}
		const kvSep, fieldSep = '=', ';'
		return ";" + encodeObject(kvSep, fieldSep, e.fields)

	default:
		panic("unreachable")
	}
}
