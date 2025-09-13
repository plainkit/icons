package lucide

import x "github.com/bloxui/blox"

// ClipboardClock creates a Clipboard Clock Lucide icon.
func ClipboardClock(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 14v2.2l1.6 1"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v.832"))),
		x.Child(x.Path(x.D("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("16"), x.R("6"))),
		x.Child(x.Rect(x.X("8"), x.Y("2"), x.RectWidth("8"), x.RectHeight("4"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
