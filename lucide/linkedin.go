package lucide

import x "github.com/bloxui/blox"

// Linkedin creates a Linkedin Lucide icon.
func Linkedin(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-linkedin", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("12"), x.X("2"), x.Y("9"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("4"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
