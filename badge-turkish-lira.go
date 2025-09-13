package lucide

import x "github.com/bloxui/blox"

// BadgeTurkishLira creates a Badge Turkish Lira Lucide icon.
func BadgeTurkishLira(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-badge-turkish-lira", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 7v10a5 5 0 0 0 5-5"))),
		x.Child(x.Path(x.D("m15 8-6 3"))),
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76"))),
	)
	return x.Svg(svgArgs...)
}
