package lucide

import x "github.com/bloxui/blox"

// Mouse creates a Mouse Lucide icon.
func Mouse(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mouse", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("20"), x.X("5"), x.Y("2"), x.Rx("7"))),
		x.Child(x.Path(x.D("M12 6v4"))),
	)
	return x.Svg(svgArgs...)
}
