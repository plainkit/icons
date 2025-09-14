package lucide

import x "github.com/bloxui/blox"

// Volleyball creates a Volleyball Lucide icon.
func Volleyball(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-volleyball", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11.1 7.1a16.55 16.55 0 0 1 10.9 4"))),
		x.Child(x.Path(x.D("M12 12a12.6 12.6 0 0 1-8.7 5"))),
		x.Child(x.Path(x.D("M16.8 13.6a16.55 16.55 0 0 1-9 7.5"))),
		x.Child(x.Path(x.D("M20.7 17a12.8 12.8 0 0 0-8.7-5 13.3 13.3 0 0 1 0-10"))),
		x.Child(x.Path(x.D("M6.3 3.8a16.55 16.55 0 0 0 1.9 11.5"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
