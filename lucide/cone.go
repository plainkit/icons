package lucide

import x "github.com/plainkit/html"

// Cone creates a Cone Lucide icon.
func Cone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m20.9 18.55-8-15.98a1 1 0 0 0-1.8 0l-8 15.98"))),
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("19"), x.EllipseRx("9"), x.EllipseRy("3"))),
	)
	return x.Svg(svgArgs...)
}
