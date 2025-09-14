package lucide

import x "github.com/bloxui/blox"

// Trash2 creates a Trash 2 Lucide icon.
func Trash2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trash-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 11v6"))),
		x.Child(x.Path(x.D("M14 11v6"))),
		x.Child(x.Path(x.D("M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"))),
		x.Child(x.Path(x.D("M3 6h18"))),
		x.Child(x.Path(x.D("M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"))),
	)
	return x.Svg(svgArgs...)
}
