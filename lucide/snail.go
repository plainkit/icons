package lucide

import x "github.com/plainkit/blox"

// Snail creates a Snail Lucide icon.
func Snail(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-snail", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 13a6 6 0 1 0 12 0 4 4 0 1 0-8 0 2 2 0 0 0 4 0"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("13"), x.R("8"))),
		x.Child(x.Path(x.D("M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6"))),
		x.Child(x.Path(x.D("M18 3 19.1 5.2"))),
		x.Child(x.Path(x.D("M22 3 20.9 5.2"))),
	)
	return x.Svg(svgArgs...)
}
