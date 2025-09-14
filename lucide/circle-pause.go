package lucide

import x "github.com/bloxui/blox"

// CirclePause creates a Circle Pause Lucide icon.
func CirclePause(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-pause", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Line(x.X1("10"), x.X2("10"), x.Y1("15"), x.Y2("9"))),
		x.Child(x.Line(x.X1("14"), x.X2("14"), x.Y1("15"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
