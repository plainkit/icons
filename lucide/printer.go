package lucide

import x "github.com/plainkit/html"

// Printer creates a Printer Lucide icon.
func Printer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-printer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"))),
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("8"), x.X("6"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
