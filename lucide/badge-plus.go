package lucide

import x "github.com/plainkit/blox"

// BadgePlus creates a Badge Plus Lucide icon.
func BadgePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("8"), x.Y2("16"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
