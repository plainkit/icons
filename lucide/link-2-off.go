package lucide

import x "github.com/plainkit/html"

// Link2Off creates a Link 2 Off Lucide icon.
func Link2Off(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-link-2-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 17H7A5 5 0 0 1 7 7"))),
		x.Child(x.Path(x.D("M15 7h2a5 5 0 0 1 4 8"))),
		x.Child(x.Line(x.X1("8"), x.X2("12"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
