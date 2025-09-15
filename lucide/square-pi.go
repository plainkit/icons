package lucide

import x "github.com/plainkit/html"

// SquarePi creates a Square Pi Lucide icon.
func SquarePi(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-pi", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 7h10"))),
		x.Child(x.Path(x.D("M10 7v10"))),
		x.Child(x.Path(x.D("M16 17a2 2 0 0 1-2-2V7"))),
	)
	return x.Svg(svgArgs...)
}
