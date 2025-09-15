package lucide

import x "github.com/plainkit/html"

// MapPinPlusInside creates a Map Pin Plus Inside Lucide icon.
func MapPinPlusInside(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-plus-inside", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"))),
		x.Child(x.Path(x.D("M12 7v6"))),
		x.Child(x.Path(x.D("M9 10h6"))),
	)
	return x.Svg(svgArgs...)
}
