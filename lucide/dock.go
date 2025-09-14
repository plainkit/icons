package lucide

import x "github.com/bloxui/blox"

// Dock creates a Dock Lucide icon.
func Dock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 8h20"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 16h12"))),
	)
	return x.Svg(svgArgs...)
}
