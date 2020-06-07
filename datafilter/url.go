package datafilter

import "net/url"

// æ˜¯æœ‰æˆçš„URL
func (obj *Object) IsURL(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, obj.err = url.ParseRequestURI(obj.rawValue)
	return obj
}
