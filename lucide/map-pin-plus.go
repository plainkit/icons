package lucide

import x "github.com/plainkit/html"

// MapPinPlus creates a Map Pin Plus Lucide icon.
func MapPinPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.914 11.105A7.298 7.298 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("M16 18h6"))),
		x.Child(x.Path(x.D("M19 15v6"))),
	)
	return x.Svg(svgArgs...)
}
