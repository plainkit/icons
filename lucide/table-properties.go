package lucide

import x "github.com/bloxui/blox"

// TableProperties creates a Table Properties Lucide icon.
func TableProperties(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table-properties", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 3v18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M21 9H3"))),
		x.Child(x.Path(x.D("M21 15H3"))),
	)
	return x.Svg(svgArgs...)
}
