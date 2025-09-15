package lucide

import x "github.com/plainkit/blox"

// AlignVerticalSpaceAround creates a Align Vertical Space Around Lucide icon.
func AlignVerticalSpaceAround(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-space-around", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("9"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 20H2"))),
		x.Child(x.Path(x.D("M22 4H2"))),
	)
	return x.Svg(svgArgs...)
}
