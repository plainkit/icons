package lucide

import x "github.com/plainkit/blox"

// SunMoon creates a Sun Moon Lucide icon.
func SunMoon(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sun-moon", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M14.837 16.385a6 6 0 1 1-7.223-7.222c.624-.147.97.66.715 1.248a4 4 0 0 0 5.26 5.259c.589-.255 1.396.09 1.248.715"))),
		x.Child(x.Path(x.D("M16 12a4 4 0 0 0-4-4"))),
		x.Child(x.Path(x.D("m19 5-1.256 1.256"))),
		x.Child(x.Path(x.D("M20 12h2"))),
	)
	return x.Svg(svgArgs...)
}
