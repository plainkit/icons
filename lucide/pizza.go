package lucide

import x "github.com/bloxui/blox"

// Pizza creates a Pizza Lucide icon.
func Pizza(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pizza", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 14-1 1"))),
		x.Child(x.Path(x.D("m13.75 18.25-1.25 1.42"))),
		x.Child(x.Path(x.D("M17.775 5.654a15.68 15.68 0 0 0-12.121 12.12"))),
		x.Child(x.Path(x.D("M18.8 9.3a1 1 0 0 0 2.1 7.7"))),
		x.Child(x.Path(x.D("M21.964 20.732a1 1 0 0 1-1.232 1.232l-18-5a1 1 0 0 1-.695-1.232A19.68 19.68 0 0 1 15.732 2.037a1 1 0 0 1 1.232.695z"))),
	)
	return x.Svg(svgArgs...)
}
