package lucide

import x "github.com/plainkit/html"

// AlignHorizontalSpaceAround creates a Align Horizontal Space Around Lucide icon.
func AlignHorizontalSpaceAround(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-space-around", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("9"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 22V2"))),
		x.Child(x.Path(x.D("M20 22V2"))),
	)
	return x.Svg(svgArgs...)
}
