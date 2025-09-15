package lucide

import x "github.com/plainkit/html"

// BriefcaseMedical creates a Briefcase Medical Lucide icon.
func BriefcaseMedical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-briefcase-medical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 11v4"))),
		x.Child(x.Path(x.D("M14 13h-4"))),
		x.Child(x.Path(x.D("M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2"))),
		x.Child(x.Path(x.D("M18 6v14"))),
		x.Child(x.Path(x.D("M6 6v14"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
