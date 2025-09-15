package lucide

import x "github.com/plainkit/html"

// Videotape creates a Videotape Lucide icon.
func Videotape(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-videotape", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 8h20"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("14"), x.R("2"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("14"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
