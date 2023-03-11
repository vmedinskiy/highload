package parser

import (
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/internal/xslices"
	"github.com/ogen-go/ogen/openapi"
	"github.com/ogen-go/ogen/uri"
)

type pathParser[P any] struct {
	// Input path.
	path string // immutable
	// Callback to lookup parameter by name.
	lookup func(name string) (P, bool) // immutable

	// Parser state.
	parts []openapi.PathPart[P] // parsed parts
	part  []rune                // current part
	param bool                  // current part is param name?
}

func pathID(path string) (string, error) {
	p, err := (&pathParser[*openapi.Parameter]{
		path: path,
		lookup: func(name string) (*openapi.Parameter, bool) {
			return nil, true
		},
	}).Parse()
	if err != nil {
		return "", err
	}
	return openapi.Path(p).ID(), nil
}

var errInvalidPathUTF8 = errors.New("path must be valid UTF-8 string")

func parsePath(path string, params []*openapi.Parameter) (openapi.Path, error) {
	if !utf8.ValidString(path) {
		return nil, errInvalidPathUTF8
	}

	// Validate and unescape path.
	//
	// FIXME(tdakkota): OpenAPI spec, as always, is not clear about path validation.
	//  All we know is that it MUST start with a slash.
	// 	At the same time, https://swagger.io/docs/specification/paths-and-operations/ says that
	// 	paths must not include query parameters.
	//  In summary, we do not pass URL scheme, user info, host or query string.
	//
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	switch {
	case u.IsAbs() || u.Host != "" || u.User != nil:
		return nil, errors.New("path MUST be relative")
	case u.RawQuery != "":
		return nil, errors.New("path MUST NOT contain a query string")
	case !strings.HasPrefix(path, "/"):
		return nil, errors.New("path MUST begin with a forward slash")
	}
	if normalized, ok := uri.NormalizeEscapedPath(path); ok {
		path = normalized
	}

	return (&pathParser[*openapi.Parameter]{
		path: path,
		lookup: func(name string) (*openapi.Parameter, bool) {
			return xslices.FindFunc(params, func(p *openapi.Parameter) bool {
				return p.Name == name && p.In == openapi.LocationPath
			})
		},
	}).Parse()
}

func parseServerURL(u string, lookup func(name string) (openapi.ServerVariable, bool)) (openapi.ServerURL, error) {
	return (&pathParser[openapi.ServerVariable]{
		path:   u,
		lookup: lookup,
	}).Parse()
}

func (p *pathParser[P]) Parse() ([]openapi.PathPart[P], error) {
	err := p.parse()
	return p.parts, err
}

func (p *pathParser[P]) parse() error {
	if !utf8.ValidString(p.path) {
		return errInvalidPathUTF8
	}
	for _, r := range p.path {
		switch r {
		case '/':
			if p.param {
				return errors.Errorf("invalid path: %s", p.path)
			}
			p.part = append(p.part, r)

		case '{':
			if p.param {
				return errors.Errorf("invalid path: %s", p.path)
			}
			if err := p.push(); err != nil {
				return err
			}
			p.param = true

		case '}':
			if !p.param {
				return errors.Errorf("invalid path: %s", p.path)
			}
			if err := p.push(); err != nil {
				return err
			}
			p.param = false

		default:
			p.part = append(p.part, r)
		}
	}

	if p.param {
		return errors.Errorf("invalid path: %s", p.path)
	}

	return p.push()
}

func (p *pathParser[P]) push() error {
	if len(p.part) == 0 {
		return nil
	}
	defer func() { p.part = nil }()

	if !p.param {
		p.parts = append(p.parts, openapi.PathPart[P]{Raw: string(p.part)})
		return nil
	}

	param, found := p.lookup(string(p.part))
	if !found {
		return errors.Errorf("path parameter not specified: %q", string(p.part))
	}

	p.parts = append(p.parts, openapi.PathPart[P]{Param: param})
	return nil
}
