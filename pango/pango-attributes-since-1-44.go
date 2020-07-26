// +build pango_1_44

package pango

// #include <pango/pango.h>
// #include "pango.go.h"
import "C"

func AttrInsertHyphensNew(insertHyphens bool) *Attribute {
	c := C.pango_attr_insert_hyphens_new(gbool(insertHyphens))
	attr := new(Attribute)
	attr.pangoAttribute = c
	return attr
}

