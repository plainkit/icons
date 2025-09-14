package lucide

import x "github.com/bloxui/blox"

// Columns2 creates a Columns 2 Lucide icon.
func Columns2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-columns-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 3v18"))),
	)
	return x.Svg(svgArgs...)
}
