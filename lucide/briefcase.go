package lucide

import x "github.com/plainkit/blox"

// Briefcase creates a Briefcase Lucide icon.
func Briefcase(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-briefcase", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
