package lucide

import x "github.com/bloxui/blox"

// Database creates a Database Lucide icon.
func Database(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-database", args)
	svgArgs = append(svgArgs,
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("5"), x.EllipseRx("9"), x.EllipseRy("3"))),
		x.Child(x.Path(x.D("M3 5V19A9 3 0 0 0 21 19V5"))),
		x.Child(x.Path(x.D("M3 12A9 3 0 0 0 21 12"))),
	)
	return x.Svg(svgArgs...)
}
