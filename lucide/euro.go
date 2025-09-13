package lucide

import x "github.com/bloxui/blox"

// Euro creates a Euro Lucide icon.
func Euro(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-euro", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 10h12"))),
		x.Child(x.Path(x.D("M4 14h9"))),
		x.Child(x.Path(x.D("M19 6a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8 2 0 3.8-.8 5.2-2"))),
	)
	return x.Svg(svgArgs...)
}
