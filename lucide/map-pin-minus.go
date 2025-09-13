package lucide

import x "github.com/bloxui/blox"

// MapPinMinus creates a Map Pin Minus Lucide icon.
func MapPinMinus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.977 14C19.6 12.701 20 11.343 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("M16 18h6"))),
	)
	return x.Svg(svgArgs...)
}
