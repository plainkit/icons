package lucide

import x "github.com/bloxui/blox"

// AlignHorizontalDistributeEnd creates a Align Horizontal Distribute End Lucide icon.
func AlignHorizontalDistributeEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-distribute-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("4"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("14"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M10 2v20"))),
		x.Child(x.Path(x.D("M20 2v20"))),
	)
	return x.Svg(svgArgs...)
}
