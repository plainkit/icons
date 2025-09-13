package lucide

import x "github.com/bloxui/blox"

// Keyboard creates a Keyboard Lucide icon.
func Keyboard(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-keyboard", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 8h.01"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
		x.Child(x.Path(x.D("M14 8h.01"))),
		x.Child(x.Path(x.D("M16 12h.01"))),
		x.Child(x.Path(x.D("M18 8h.01"))),
		x.Child(x.Path(x.D("M6 8h.01"))),
		x.Child(x.Path(x.D("M7 16h10"))),
		x.Child(x.Path(x.D("M8 12h.01"))),
		x.Child(x.Rect(x.X("2"), x.Y("4"), x.RectWidth("20"), x.RectHeight("16"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
