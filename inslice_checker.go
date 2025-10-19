// Copyright (c) 2025 by Gilbert Ramirez <gram@alumni.rice.edu>.
// All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gocheck_extra

import (
	"reflect"

	. "gopkg.in/check.v1"
)

type inSliceChecker struct {
	info *CheckerInfo
}

func (s *inSliceChecker) Info() *CheckerInfo {
	return s.info
}

func (s *inSliceChecker) Check(params []any, names []string) (result bool, error string) {

	if len(params) != 2 {
		panic("Need 2 params")
	}
	item := params[0]
	itemType := reflect.TypeOf(item)
	itemValue := reflect.ValueOf(item)

	containerAny := params[1]
	containerValue := reflect.ValueOf(containerAny)
	if containerValue.Kind() != reflect.Slice {
		return false, "2nd argument isn't a slice"
	}

	// Is the slice convertable to the right type?
	if !containerValue.CanConvert(reflect.SliceOf(itemType)) {
		return false, "the slice doesn't contain the same type of items"
	}

	for i := 0; i < containerValue.Len(); i++ {
		elementValue := containerValue.Index(i)
		if elementValue.Equal(itemValue) {
			return true, ""
		}
	}

	return false, "Item is not in the slice"
}

// Checks if an item is in a slice of the same type
// Usage:
// c.Check(item, InSlice, slice)
var InSlice Checker = &inSliceChecker{
	&CheckerInfo{Name: "InSlice", Params: []string{"item", "slice"}},
}
