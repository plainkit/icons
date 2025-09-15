package lucide

import x "github.com/plainkit/html"

// MousePointerClick creates a Mouse Pointer Click Lucide icon.
func MousePointerClick(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mouse-pointer-click", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 4.1 12 6"))),
		x.Child(x.Path(x.D("m5.1 8-2.9-.8"))),
		x.Child(x.Path(x.D("m6 12-1.9 2"))),
		x.Child(x.Path(x.D("M7.2 2.2 8 5.1"))),
		x.Child(x.Path(x.D("M9.037 9.69a.498.498 0 0 1 .653-.653l11 4.5a.5.5 0 0 1-.074.949l-4.349 1.041a1 1 0 0 0-.74.739l-1.04 4.35a.5.5 0 0 1-.95.074z"))),
	)
	return x.Svg(svgArgs...)
}
