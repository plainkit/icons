package lucide

import x "github.com/bloxui/blox"

// Link creates a Link Lucide icon.
func Link(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-link", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"))),
		x.Child(x.Path(x.D("M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"))),
	)
	return x.Svg(svgArgs...)
}
