package lucide

import x "github.com/plainkit/blox"

// Wine creates a Wine Lucide icon.
func Wine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wine", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 22h8"))),
		x.Child(x.Path(x.D("M7 10h10"))),
		x.Child(x.Path(x.D("M12 15v7"))),
		x.Child(x.Path(x.D("M12 15a5 5 0 0 0 5-5c0-2-.5-4-2-8H9c-1.5 4-2 6-2 8a5 5 0 0 0 5 5Z"))),
	)
	return x.Svg(svgArgs...)
}
