package lucide

import x "github.com/plainkit/html"

// CircleSlash creates a Circle Slash Lucide icon.
func CircleSlash(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-slash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Line(x.X1("9"), x.X2("15"), x.Y1("15"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
