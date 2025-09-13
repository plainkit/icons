package lucide

import x "github.com/bloxui/blox"

// BriefcaseBusiness creates a Briefcase Business Lucide icon.
func BriefcaseBusiness(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-briefcase-business", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12h.01"))),
		x.Child(x.Path(x.D("M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2"))),
		x.Child(x.Path(x.D("M22 13a18.15 18.15 0 0 1-20 0"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
