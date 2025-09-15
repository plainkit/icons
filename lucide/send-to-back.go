package lucide

import x "github.com/plainkit/html"

// SendToBack creates a Send To Back Lucide icon.
func SendToBack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-send-to-back", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("14"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("2"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 14v1a2 2 0 0 0 2 2h1"))),
		x.Child(x.Path(x.D("M14 7h1a2 2 0 0 1 2 2v1"))),
	)
	return x.Svg(svgArgs...)
}
