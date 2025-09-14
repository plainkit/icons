package lucide

import x "github.com/bloxui/blox"

// BaggageClaim creates a Baggage Claim Lucide icon.
func BaggageClaim(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-baggage-claim", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2"))),
		x.Child(x.Path(x.D("M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10"))),
		x.Child(x.Rect(x.RectWidth("13"), x.RectHeight("8"), x.X("8"), x.Y("6"), x.Rx("1"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("20"), x.R("2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("20"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
