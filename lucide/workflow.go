package lucide

import x "github.com/bloxui/blox"

// Workflow creates a Workflow Lucide icon.
func Workflow(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-workflow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 11v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("13"), x.Y("13"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
