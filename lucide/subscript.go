package lucide

import x "github.com/bloxui/blox"

// Subscript creates a Subscript Lucide icon.
func Subscript(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-subscript", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m4 5 8 8"))),
		x.Child(x.Path(x.D("m12 5-8 8"))),
		x.Child(x.Path(x.D("M20 19h-4c0-1.5.44-2 1.5-2.5S20 15.33 20 14c0-.47-.17-.93-.48-1.29a2.11 2.11 0 0 0-2.62-.44c-.42.24-.74.62-.9 1.07"))),
	)
	return x.Svg(svgArgs...)
}
