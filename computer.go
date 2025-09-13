package lucide

import x "github.com/bloxui/blox"

// Computer creates a Computer Lucide icon.
func Computer(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-computer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("8"), x.X("5"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("8"), x.X("2"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 18h2"))),
		x.Child(x.Path(x.D("M12 18h6"))),
	)
	return x.Svg(svgArgs...)
}