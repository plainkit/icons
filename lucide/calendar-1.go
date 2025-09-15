package lucide

import x "github.com/plainkit/html"

// Calendar1 creates a Calendar 1 Lucide icon.
func Calendar1(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 14h1v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
