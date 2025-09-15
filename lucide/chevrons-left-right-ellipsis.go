package lucide

import x "github.com/plainkit/blox"

// ChevronsLeftRightEllipsis creates a Chevrons Left Right Ellipsis Lucide icon.
func ChevronsLeftRightEllipsis(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-left-right-ellipsis", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12h.01"))),
		x.Child(x.Path(x.D("M16 12h.01"))),
		x.Child(x.Path(x.D("m17 7 5 5-5 5"))),
		x.Child(x.Path(x.D("m7 7-5 5 5 5"))),
		x.Child(x.Path(x.D("M8 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
