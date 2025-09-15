package lucide

import x "github.com/plainkit/blox"

// Bomb creates a Bomb Lucide icon.
func Bomb(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bomb", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11"), x.Cy("13"), x.R("9"))),
		x.Child(x.Path(x.D("M14.35 4.65 16.3 2.7a2.41 2.41 0 0 1 3.4 0l1.6 1.6a2.4 2.4 0 0 1 0 3.4l-1.95 1.95"))),
		x.Child(x.Path(x.D("m22 2-1.5 1.5"))),
	)
	return x.Svg(svgArgs...)
}
