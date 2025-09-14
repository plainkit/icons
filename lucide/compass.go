package lucide

import x "github.com/bloxui/blox"

// Compass creates a Compass Lucide icon.
func Compass(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-compass", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16.24 7.76-1.804 5.411a2 2 0 0 1-1.265 1.265L7.76 16.24l1.804-5.411a2 2 0 0 1 1.265-1.265z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
