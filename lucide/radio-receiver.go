package lucide

import x "github.com/plainkit/html"

// RadioReceiver creates a Radio Receiver Lucide icon.
func RadioReceiver(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-radio-receiver", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 16v2"))),
		x.Child(x.Path(x.D("M19 16v2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("8"), x.X("2"), x.Y("8"), x.Rx("2"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
