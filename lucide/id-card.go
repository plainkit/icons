package lucide

import x "github.com/bloxui/blox"

// IdCard creates a Id Card Lucide icon.
func IdCard(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-id-card", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 10h2"))),
		x.Child(x.Path(x.D("M16 14h2"))),
		x.Child(x.Path(x.D("M6.17 15a3 3 0 0 1 5.66 0"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("11"), x.R("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
