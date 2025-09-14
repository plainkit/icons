package lucide

import x "github.com/bloxui/blox"

// Dessert creates a Dessert Lucide icon.
func Dessert(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dessert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.162 3.167A10 10 0 0 0 2 13a2 2 0 0 0 4 0v-1a2 2 0 0 1 4 0v4a2 2 0 0 0 4 0v-4a2 2 0 0 1 4 0v1a2 2 0 0 0 4-.006 10 10 0 0 0-8.161-9.826"))),
		x.Child(x.Path(x.D("M20.804 14.869a9 9 0 0 1-17.608 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("4"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
