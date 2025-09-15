package lucide

import x "github.com/plainkit/blox"

// Anchor creates a Anchor Lucide icon.
func Anchor(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-anchor", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22V8"))),
		x.Child(x.Path(x.D("M5 12H2a10 10 0 0 0 20 0h-3"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
