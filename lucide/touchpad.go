package lucide

import x "github.com/bloxui/blox"

// Touchpad creates a Touchpad Lucide icon.
func Touchpad(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-touchpad", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 14h20"))),
		x.Child(x.Path(x.D("M12 20v-6"))),
	)
	return x.Svg(svgArgs...)
}
