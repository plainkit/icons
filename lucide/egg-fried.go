package lucide

import x "github.com/bloxui/blox"

// EggFried creates a Egg Fried Lucide icon.
func EggFried(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-egg-fried", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11.5"), x.Cy("12.5"), x.R("3.5"))),
		x.Child(x.Path(x.D("M3 8c0-3.5 2.5-6 6.5-6 5 0 4.83 3 7.5 5s5 2 5 6c0 4.5-2.5 6.5-7 6.5-2.5 0-2.5 2.5-6 2.5s-7-2-7-5.5c0-3 1.5-3 1.5-5C3.5 10 3 9 3 8Z"))),
	)
	return x.Svg(svgArgs...)
}
