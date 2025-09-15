package lucide

import x "github.com/plainkit/html"

// Battery creates a Battery Lucide icon.
func Battery(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-battery", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M 22 14 L 22 10"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
