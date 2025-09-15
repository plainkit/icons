package lucide

import x "github.com/plainkit/blox"

// Drill creates a Drill Lucide icon.
func Drill(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-drill", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 18a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a3 3 0 0 1-3-3 1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M13 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1l-.81 3.242a1 1 0 0 1-.97.758H8"))),
		x.Child(x.Path(x.D("M14 4h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3"))),
		x.Child(x.Path(x.D("M18 6h4"))),
		x.Child(x.Path(x.D("m5 10-2 8"))),
		x.Child(x.Path(x.D("m7 18 2-8"))),
	)
	return x.Svg(svgArgs...)
}
