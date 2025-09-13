package lucide

import x "github.com/bloxui/blox"

// BriefcaseConveyorBelt creates a Briefcase Conveyor Belt Lucide icon.
func BriefcaseConveyorBelt(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-briefcase-conveyor-belt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 20v2"))),
		x.Child(x.Path(x.D("M14 20v2"))),
		x.Child(x.Path(x.D("M18 20v2"))),
		x.Child(x.Path(x.D("M21 20H3"))),
		x.Child(x.Path(x.D("M6 20v2"))),
		x.Child(x.Path(x.D("M8 16V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v12"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("10"), x.X("4"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
