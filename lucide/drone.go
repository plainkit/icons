package lucide

import x "github.com/bloxui/blox"

// Drone creates a Drone Lucide icon.
func Drone(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-drone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10 7 7"))),
		x.Child(x.Path(x.D("m10 14-3 3"))),
		x.Child(x.Path(x.D("m14 10 3-3"))),
		x.Child(x.Path(x.D("m14 14 3 3"))),
		x.Child(x.Path(x.D("M14.205 4.139a4 4 0 1 1 5.439 5.863"))),
		x.Child(x.Path(x.D("M19.637 14a4 4 0 1 1-5.432 5.868"))),
		x.Child(x.Path(x.D("M4.367 10a4 4 0 1 1 5.438-5.862"))),
		x.Child(x.Path(x.D("M9.795 19.862a4 4 0 1 1-5.429-5.873"))),
		x.Child(x.Rect(x.X("10"), x.Y("8"), x.RectWidth("4"), x.RectHeight("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
