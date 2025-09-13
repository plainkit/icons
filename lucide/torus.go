package lucide

import x "github.com/bloxui/blox"

// Torus creates a Torus Lucide icon.
func Torus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-torus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("11"), x.EllipseRx("3"), x.EllipseRy("2"))),
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("12.5"), x.EllipseRx("10"), x.EllipseRy("8.5"))),
	)
	return x.Svg(svgArgs...)
}
