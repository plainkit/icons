package lucide

import x "github.com/bloxui/blox"

// BadgeX creates a Badge X Lucide icon.
func BadgeX(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-badge-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Line(x.X1("15"), x.X2("9"), x.Y1("9"), x.Y2("15"))),
		x.Child(x.Line(x.X1("9"), x.X2("15"), x.Y1("9"), x.Y2("15"))),
	)
	return x.Svg(svgArgs...)
}
