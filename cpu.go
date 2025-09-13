package lucide

import x "github.com/bloxui/blox"

// Cpu creates a Cpu Lucide icon.
func Cpu(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cpu", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M17 20v2"))),
		x.Child(x.Path(x.D("M17 2v2"))),
		x.Child(x.Path(x.D("M2 12h2"))),
		x.Child(x.Path(x.D("M2 17h2"))),
		x.Child(x.Path(x.D("M2 7h2"))),
		x.Child(x.Path(x.D("M20 12h2"))),
		x.Child(x.Path(x.D("M20 17h2"))),
		x.Child(x.Path(x.D("M20 7h2"))),
		x.Child(x.Path(x.D("M7 20v2"))),
		x.Child(x.Path(x.D("M7 2v2"))),
		x.Child(x.Rect(x.X("4"), x.Y("4"), x.RectWidth("16"), x.RectHeight("16"), x.Rx("2"))),
		x.Child(x.Rect(x.X("8"), x.Y("8"), x.RectWidth("8"), x.RectHeight("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}