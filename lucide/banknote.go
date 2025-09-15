package lucide

import x "github.com/plainkit/blox"

// Banknote creates a Banknote Lucide icon.
func Banknote(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-banknote", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("M6 12h.01M18 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
