package lucide

import x "github.com/bloxui/blox"

// Construction creates a Construction Lucide icon.
func Construction(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-construction", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.X("2"), x.Y("6"), x.RectWidth("20"), x.RectHeight("8"), x.Rx("1"))),
		x.Child(x.Path(x.D("M17 14v7"))),
		x.Child(x.Path(x.D("M7 14v7"))),
		x.Child(x.Path(x.D("M17 3v3"))),
		x.Child(x.Path(x.D("M7 3v3"))),
		x.Child(x.Path(x.D("M10 14 2.3 6.3"))),
		x.Child(x.Path(x.D("m14 6 7.7 7.7"))),
		x.Child(x.Path(x.D("m8 6 8 8"))),
	)
	return x.Svg(svgArgs...)
}
