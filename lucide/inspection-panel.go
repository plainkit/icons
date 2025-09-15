package lucide

import x "github.com/plainkit/blox"

// InspectionPanel creates a Inspection Panel Lucide icon.
func InspectionPanel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-inspection-panel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 7h.01"))),
		x.Child(x.Path(x.D("M17 7h.01"))),
		x.Child(x.Path(x.D("M7 17h.01"))),
		x.Child(x.Path(x.D("M17 17h.01"))),
	)
	return x.Svg(svgArgs...)
}
