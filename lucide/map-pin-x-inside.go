package lucide

import x "github.com/plainkit/html"

// MapPinXInside creates a Map Pin X Inside Lucide icon.
func MapPinXInside(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-x-inside", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"))),
		x.Child(x.Path(x.D("m14.5 7.5-5 5"))),
		x.Child(x.Path(x.D("m9.5 7.5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
