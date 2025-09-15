package lucide

import x "github.com/plainkit/html"

// CupSoda creates a Cup Soda Lucide icon.
func CupSoda(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cup-soda", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 8 1.75 12.28a2 2 0 0 0 2 1.72h4.54a2 2 0 0 0 2-1.72L18 8"))),
		x.Child(x.Path(x.D("M5 8h14"))),
		x.Child(x.Path(x.D("M7 15a6.47 6.47 0 0 1 5 0 6.47 6.47 0 0 0 5 0"))),
		x.Child(x.Path(x.D("m12 8 1-6h2"))),
	)
	return x.Svg(svgArgs...)
}
