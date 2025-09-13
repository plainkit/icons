package lucide

import x "github.com/bloxui/blox"

// Haze creates a Haze Lucide icon.
func Haze(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-haze", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m5.2 6.2 1.4 1.4"))),
		x.Child(x.Path(x.D("M2 13h2"))),
		x.Child(x.Path(x.D("M20 13h2"))),
		x.Child(x.Path(x.D("m17.4 7.6 1.4-1.4"))),
		x.Child(x.Path(x.D("M22 17H2"))),
		x.Child(x.Path(x.D("M22 21H2"))),
		x.Child(x.Path(x.D("M16 13a4 4 0 0 0-8 0"))),
		x.Child(x.Path(x.D("M12 5V2.5"))),
	)
	return x.Svg(svgArgs...)
}
