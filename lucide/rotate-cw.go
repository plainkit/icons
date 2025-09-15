package lucide

import x "github.com/plainkit/html"

// RotateCw creates a Rotate Cw Lucide icon.
func RotateCw(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rotate-cw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"))),
		x.Child(x.Path(x.D("M21 3v5h-5"))),
	)
	return x.Svg(svgArgs...)
}
