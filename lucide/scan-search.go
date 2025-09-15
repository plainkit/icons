package lucide

import x "github.com/plainkit/blox"

// ScanSearch creates a Scan Search Lucide icon.
func ScanSearch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scan-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("m16 16-1.9-1.9"))),
	)
	return x.Svg(svgArgs...)
}
