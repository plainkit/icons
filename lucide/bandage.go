package lucide

import x "github.com/bloxui/blox"

// Bandage creates a Bandage Lucide icon.
func Bandage(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bandage", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10.01h.01"))),
		x.Child(x.Path(x.D("M10 14.01h.01"))),
		x.Child(x.Path(x.D("M14 10.01h.01"))),
		x.Child(x.Path(x.D("M14 14.01h.01"))),
		x.Child(x.Path(x.D("M18 6v11.5"))),
		x.Child(x.Path(x.D("M6 6v12"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
