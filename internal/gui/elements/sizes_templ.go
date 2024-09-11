// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package elements

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type Size int

const (
	SizeDefault Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
)

func (s Size) CardAttributes() templ.Attributes {
	switch s {
	case SizeSmall:
		return templ.Attributes{
			"class": "max-w-md bg-white border rounded-lg shadow-sm p-7 border-neutral-200/60",
		}
	case SizeLarge:
		return templ.Attributes{
			"class": "max-w-xl bg-white border rounded-lg shadow-sm p-7 border-neutral-200/60",
		}
	}
	return templ.Attributes{
		"class": "max-w-lg bg-white border rounded-lg shadow-sm p-7 border-neutral-200/60",
	}
}

func (s Size) SvgAttributes() templ.Attributes {
	switch s {
	case SizeSmall:
		return templ.Attributes{
			"height": "16",
			"width":  "16",
		}
	case SizeLarge:
		return templ.Attributes{
			"height": "32",
			"width":  "32",
		}
	}
	return templ.Attributes{
		"height": "24",
		"width":  "24",
	}
}

func (s Size) TextAttributes() templ.Attributes {
	switch s {
	case SizeSmall:
		return templ.Attributes{
			"class": "text-sm",
		}
	case SizeLarge:
		return templ.Attributes{
			"class": "text-lg",
		}
	}
	return templ.Attributes{
		"class": "text-md",
	}
}

var _ = templruntime.GeneratedTemplate
