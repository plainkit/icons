package lucide

import x "github.com/bloxui/blox"

// TramFront creates a Tram Front Lucide icon.
func TramFront(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-tram-front", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("16"), x.X("4"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 11h16"))),
		x.Child(x.Path(x.D("M12 3v8"))),
		x.Child(x.Path(x.D("m8 19-2 3"))),
		x.Child(x.Path(x.D("m18 22-2-3"))),
		x.Child(x.Path(x.D("M8 15h.01"))),
		x.Child(x.Path(x.D("M16 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
