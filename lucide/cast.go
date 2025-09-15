package lucide

import x "github.com/plainkit/blox"

// Cast creates a Cast Lucide icon.
func Cast(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cast", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6"))),
		x.Child(x.Path(x.D("M2 12a9 9 0 0 1 8 8"))),
		x.Child(x.Path(x.D("M2 16a5 5 0 0 1 4 4"))),
		x.Child(x.Line(x.X1("2"), x.X2("2.01"), x.Y1("20"), x.Y2("20"))),
	)
	return x.Svg(svgArgs...)
}
