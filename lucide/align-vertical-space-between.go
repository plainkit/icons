package lucide

import x "github.com/plainkit/blox"

// AlignVerticalSpaceBetween creates a Align Vertical Space Between Lucide icon.
func AlignVerticalSpaceBetween(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-space-between", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("6"), x.X("5"), x.Y("15"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 21h20"))),
		x.Child(x.Path(x.D("M2 3h20"))),
	)
	return x.Svg(svgArgs...)
}
