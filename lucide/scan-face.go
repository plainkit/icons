package lucide

import x "github.com/plainkit/html"

// ScanFace creates a Scan Face Lucide icon.
func ScanFace(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scan-face", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Path(x.D("M8 14s1.5 2 4 2 4-2 4-2"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Path(x.D("M15 9h.01"))),
	)
	return x.Svg(svgArgs...)
}
