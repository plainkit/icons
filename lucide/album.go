package lucide

import x "github.com/bloxui/blox"

// Album creates a Album Lucide icon.
func Album(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-album", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Polyline(x.Points("11 3 11 11 14 8 17 11 17 3"))),
	)
	return x.Svg(svgArgs...)
}
