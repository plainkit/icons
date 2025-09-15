package lucide

import x "github.com/plainkit/blox"

// CardSim creates a Card Sim Lucide icon.
func CardSim(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-card-sim", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 14v4"))),
		x.Child(x.Path(x.D("M14.172 2a2 2 0 0 1 1.414.586l3.828 3.828A2 2 0 0 1 20 7.828V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("M8 14h8"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("8"), x.Y("10"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
