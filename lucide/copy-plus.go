package lucide

import x "github.com/bloxui/blox"

// CopyPlus creates a Copy Plus Lucide icon.
func CopyPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-copy-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("15"), x.X2("15"), x.Y1("12"), x.Y2("18"))),
		x.Child(x.Line(x.X1("12"), x.X2("18"), x.Y1("15"), x.Y2("15"))),
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("14"), x.X("8"), x.Y("8"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"))),
	)
	return x.Svg(svgArgs...)
}
