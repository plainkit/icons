package lucide

import x "github.com/plainkit/blox"

// Slack creates a Slack Lucide icon.
func Slack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-slack", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("3"), x.RectHeight("8"), x.X("13"), x.Y("2"), x.Rx("1.5"))),
		x.Child(x.Path(x.D("M19 8.5V10h1.5A1.5 1.5 0 1 0 19 8.5"))),
		x.Child(x.Rect(x.RectWidth("3"), x.RectHeight("8"), x.X("8"), x.Y("14"), x.Rx("1.5"))),
		x.Child(x.Path(x.D("M5 15.5V14H3.5A1.5 1.5 0 1 0 5 15.5"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("3"), x.X("14"), x.Y("13"), x.Rx("1.5"))),
		x.Child(x.Path(x.D("M15.5 19H14v1.5a1.5 1.5 0 1 0 1.5-1.5"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("3"), x.X("2"), x.Y("8"), x.Rx("1.5"))),
		x.Child(x.Path(x.D("M8.5 5H10V3.5A1.5 1.5 0 1 0 8.5 5"))),
	)
	return x.Svg(svgArgs...)
}
