package lucide

import x "github.com/bloxui/blox"

// PanelsTopLeft creates a Panels Top Left Lucide icon.
func PanelsTopLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-panels-top-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("M9 21V9"))),
	)
	return x.Svg(svgArgs...)
}
