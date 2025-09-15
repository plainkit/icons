package lucide

import x "github.com/plainkit/blox"

// FileSearch creates a File Search Lucide icon.
func FileSearch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M4.268 21a2 2 0 0 0 1.727 1H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("m9 18-1.5-1.5"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("14"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
