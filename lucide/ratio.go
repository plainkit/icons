package lucide

import x "github.com/plainkit/blox"

// Ratio creates a Ratio Lucide icon.
func Ratio(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ratio", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("20"), x.X("6"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
