package lucide

import x "github.com/bloxui/blox"

// Amphora creates a Amphora Lucide icon.
func Amphora(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-amphora", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v5.632c0 .424-.272.795-.653.982A6 6 0 0 0 6 14c.006 4 3 7 5 8"))),
		x.Child(x.Path(x.D("M10 5H8a2 2 0 0 0 0 4h.68"))),
		x.Child(x.Path(x.D("M14 2v5.632c0 .424.272.795.652.982A6 6 0 0 1 18 14c0 4-3 7-5 8"))),
		x.Child(x.Path(x.D("M14 5h2a2 2 0 0 1 0 4h-.68"))),
		x.Child(x.Path(x.D("M18 22H6"))),
		x.Child(x.Path(x.D("M9 2h6"))),
	)
	return x.Svg(svgArgs...)
}
