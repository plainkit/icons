package lucide

import x "github.com/plainkit/blox"

// Square creates a Square Lucide icon.
func Square(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
