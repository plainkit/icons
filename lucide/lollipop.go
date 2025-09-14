package lucide

import x "github.com/bloxui/blox"

// Lollipop creates a Lollipop Lucide icon.
func Lollipop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lollipop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
		x.Child(x.Path(x.D("m21 21-4.3-4.3"))),
		x.Child(x.Path(x.D("M11 11a2 2 0 0 0 4 0 4 4 0 0 0-8 0 6 6 0 0 0 12 0"))),
	)
	return x.Svg(svgArgs...)
}
