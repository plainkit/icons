package lucide

import x "github.com/bloxui/blox"

// Drum creates a Drum Lucide icon.
func Drum(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-drum", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 2 8 8"))),
		x.Child(x.Path(x.D("m22 2-8 8"))),
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("9"), x.EllipseRx("10"), x.EllipseRy("5"))),
		x.Child(x.Path(x.D("M7 13.4v7.9"))),
		x.Child(x.Path(x.D("M12 14v8"))),
		x.Child(x.Path(x.D("M17 13.4v7.9"))),
		x.Child(x.Path(x.D("M2 9v8a10 5 0 0 0 20 0V9"))),
	)
	return x.Svg(svgArgs...)
}
