package lucide

import x "github.com/plainkit/blox"

// Trees creates a Trees Lucide icon.
func Trees(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trees", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10v.2A3 3 0 0 1 8.9 16H5a3 3 0 0 1-1-5.8V10a3 3 0 0 1 6 0Z"))),
		x.Child(x.Path(x.D("M7 16v6"))),
		x.Child(x.Path(x.D("M13 19v3"))),
		x.Child(x.Path(x.D("M12 19h8.3a1 1 0 0 0 .7-1.7L18 14h.3a1 1 0 0 0 .7-1.7L16 9h.2a1 1 0 0 0 .8-1.7L13 3l-1.4 1.5"))),
	)
	return x.Svg(svgArgs...)
}
