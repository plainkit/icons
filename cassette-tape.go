package lucide

import x "github.com/bloxui/blox"

// CassetteTape creates a Cassette Tape Lucide icon.
func CassetteTape(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cassette-tape", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("10"), x.R("2"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("10"), x.R("2"))),
		x.Child(x.Path(x.D("m6 20 .7-2.9A1.4 1.4 0 0 1 8.1 16h7.8a1.4 1.4 0 0 1 1.4 1l.7 3"))),
	)
	return x.Svg(svgArgs...)
}
