package lucide

import x "github.com/plainkit/html"

// ScanText creates a Scan Text Lucide icon.
func ScanText(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scan-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Path(x.D("M7 8h8"))),
		x.Child(x.Path(x.D("M7 12h10"))),
		x.Child(x.Path(x.D("M7 16h6"))),
	)
	return x.Svg(svgArgs...)
}
