package lucide

import x "github.com/bloxui/blox"

// FoldHorizontal creates a Fold Horizontal Lucide icon.
func FoldHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fold-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 12h6"))),
		x.Child(x.Path(x.D("M22 12h-6"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M12 8v2"))),
		x.Child(x.Path(x.D("M12 14v2"))),
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("m19 9-3 3 3 3"))),
		x.Child(x.Path(x.D("m5 15 3-3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
