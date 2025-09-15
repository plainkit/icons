package lucide

import x "github.com/plainkit/html"

// SmilePlus creates a Smile Plus Lucide icon.
func SmilePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-smile-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 11v1a10 10 0 1 1-9-10"))),
		x.Child(x.Path(x.D("M8 14s1.5 2 4 2 4-2 4-2"))),
		x.Child(x.Line(x.X1("9"), x.X2("9.01"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Path(x.D("M16 5h6"))),
		x.Child(x.Path(x.D("M19 2v6"))),
	)
	return x.Svg(svgArgs...)
}
