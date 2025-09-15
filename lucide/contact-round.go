package lucide

import x "github.com/plainkit/blox"

// ContactRound creates a Contact Round Lucide icon.
func ContactRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-contact-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 2v2"))),
		x.Child(x.Path(x.D("M17.915 22a6 6 0 0 0-12 0"))),
		x.Child(x.Path(x.D("M8 2v2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
