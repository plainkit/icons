package lucide

import x "github.com/plainkit/html"

// AlignCenterHorizontal creates a Align Center Horizontal Lucide icon.
func AlignCenterHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-center-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("M10 16v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4"))),
		x.Child(x.Path(x.D("M10 8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M20 16v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1"))),
		x.Child(x.Path(x.D("M14 8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1"))),
	)
	return x.Svg(svgArgs...)
}
