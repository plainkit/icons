package lucide

import x "github.com/plainkit/blox"

// MapPinPen creates a Map Pin Pen Lucide icon.
func MapPinPen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-pen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.97 9.304A8 8 0 0 0 2 10c0 4.69 4.887 9.562 7.022 11.468"))),
		x.Child(x.Path(x.D("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("10"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
