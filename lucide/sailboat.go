package lucide

import x "github.com/plainkit/html"

// Sailboat creates a Sailboat Lucide icon.
func Sailboat(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sailboat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v15"))),
		x.Child(x.Path(x.D("M7 22a4 4 0 0 1-4-4 1 1 0 0 1 1-1h16a1 1 0 0 1 1 1 4 4 0 0 1-4 4z"))),
		x.Child(x.Path(x.D("M9.159 2.46a1 1 0 0 1 1.521-.193l9.977 8.98A1 1 0 0 1 20 13H4a1 1 0 0 1-.824-1.567z"))),
	)
	return x.Svg(svgArgs...)
}
