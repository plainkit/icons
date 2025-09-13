package lucide

import x "github.com/bloxui/blox"

// Contact creates a Contact Lucide icon.
func Contact(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-contact", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 2v2"))),
		x.Child(x.Path(x.D("M7 22v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M8 2v2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("11"), x.R("3"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
