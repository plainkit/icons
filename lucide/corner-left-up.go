package lucide

import x "github.com/plainkit/html"

// CornerLeftUp creates a Corner Left Up Lucide icon.
func CornerLeftUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-left-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 9 9 4 4 9"))),
		x.Child(x.Path(x.D("M20 20h-7a4 4 0 0 1-4-4V4"))),
	)
	return x.Svg(svgArgs...)
}
