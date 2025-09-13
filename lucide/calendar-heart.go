package lucide

import x "github.com/bloxui/blox"

// CalendarHeart creates a Calendar Heart Lucide icon.
func CalendarHeart(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-heart", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.127 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v5.125"))),
		x.Child(x.Path(x.D("M14.62 18.8A2.25 2.25 0 1 1 18 15.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
