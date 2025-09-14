package lucide

import x "github.com/bloxui/blox"

// Link2 creates a Link 2 Lucide icon.
func Link2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-link-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 17H7A5 5 0 0 1 7 7h2"))),
		x.Child(x.Path(x.D("M15 7h2a5 5 0 1 1 0 10h-2"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
