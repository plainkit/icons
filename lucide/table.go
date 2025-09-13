package lucide

import x "github.com/bloxui/blox"

// Table creates a Table Lucide icon.
func Table(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("M3 15h18"))),
	)
	return x.Svg(svgArgs...)
}
