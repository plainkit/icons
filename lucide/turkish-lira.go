package lucide

import x "github.com/bloxui/blox"

// TurkishLira creates a Turkish Lira Lucide icon.
func TurkishLira(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-turkish-lira", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 4 5 9"))),
		x.Child(x.Path(x.D("m15 8.5-10 5"))),
		x.Child(x.Path(x.D("M18 12a9 9 0 0 1-9 9V3"))),
	)
	return x.Svg(svgArgs...)
}
