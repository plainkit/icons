package lucide

import x "github.com/plainkit/html"

// Files creates a Files Lucide icon.
func Files(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-files", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2a2 2 0 0 1 1.414.586l4 4A2 2 0 0 1 21 8v7a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("M15 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M5 7a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h8a2 2 0 0 0 1.732-1"))),
	)
	return x.Svg(svgArgs...)
}
