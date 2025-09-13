package lucide

import x "github.com/bloxui/blox"

// Hotel creates a Hotel Lucide icon.
func Hotel(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hotel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 22v-6.57"))),
		x.Child(x.Path(x.D("M12 11h.01"))),
		x.Child(x.Path(x.D("M12 7h.01"))),
		x.Child(x.Path(x.D("M14 15.43V22"))),
		x.Child(x.Path(x.D("M15 16a5 5 0 0 0-6 0"))),
		x.Child(x.Path(x.D("M16 11h.01"))),
		x.Child(x.Path(x.D("M16 7h.01"))),
		x.Child(x.Path(x.D("M8 11h.01"))),
		x.Child(x.Path(x.D("M8 7h.01"))),
		x.Child(x.Rect(x.X("4"), x.Y("2"), x.RectWidth("16"), x.RectHeight("20"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
