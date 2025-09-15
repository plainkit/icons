package lucide

import x "github.com/plainkit/html"

// Section creates a Section Lucide icon.
func Section(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-section", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5a4 3 0 0 0-8 0c0 4 8 3 8 7a4 3 0 0 1-8 0"))),
		x.Child(x.Path(x.D("M8 19a4 3 0 0 0 8 0c0-4-8-3-8-7a4 3 0 0 1 8 0"))),
	)
	return x.Svg(svgArgs...)
}
