package lucide

import x "github.com/plainkit/blox"

// Cylinder creates a Cylinder Lucide icon.
func Cylinder(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cylinder", args)
	svgArgs = append(svgArgs,
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("5"), x.EllipseRx("9"), x.EllipseRy("3"))),
		x.Child(x.Path(x.D("M3 5v14a9 3 0 0 0 18 0V5"))),
	)
	return x.Svg(svgArgs...)
}
