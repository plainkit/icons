package lucide

import x "github.com/bloxui/blox"

// Columns3 creates a Columns 3 Lucide icon.
func Columns3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-columns-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Path(x.D("M15 3v18"))),
	)
	return x.Svg(svgArgs...)
}
