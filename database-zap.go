package lucide

import x "github.com/bloxui/blox"

// DatabaseZap creates a Database Zap Lucide icon.
func DatabaseZap(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-database-zap", args)
	svgArgs = append(svgArgs,
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("5"), x.EllipseRx("9"), x.EllipseRy("3"))),
		x.Child(x.Path(x.D("M3 5V19A9 3 0 0 0 15 21.84"))),
		x.Child(x.Path(x.D("M21 5V8"))),
		x.Child(x.Path(x.D("M21 12L18 17H22L19 22"))),
		x.Child(x.Path(x.D("M3 12A9 3 0 0 0 14.59 14.87"))),
	)
	return x.Svg(svgArgs...)
}