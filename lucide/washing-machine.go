package lucide

import x "github.com/bloxui/blox"

// WashingMachine creates a Washing Machine Lucide icon.
func WashingMachine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-washing-machine", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 6h3"))),
		x.Child(x.Path(x.D("M17 6h.01"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("20"), x.X("3"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("5"))),
		x.Child(x.Path(x.D("M12 18a2.5 2.5 0 0 0 0-5 2.5 2.5 0 0 1 0-5"))),
	)
	return x.Svg(svgArgs...)
}
