package lucide

import x "github.com/bloxui/blox"

// Film creates a Film Lucide icon.
func Film(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-film", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 3v18"))),
		x.Child(x.Path(x.D("M3 7.5h4"))),
		x.Child(x.Path(x.D("M3 12h18"))),
		x.Child(x.Path(x.D("M3 16.5h4"))),
		x.Child(x.Path(x.D("M17 3v18"))),
		x.Child(x.Path(x.D("M17 7.5h4"))),
		x.Child(x.Path(x.D("M17 16.5h4"))),
	)
	return x.Svg(svgArgs...)
}
