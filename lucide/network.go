package lucide

import x "github.com/plainkit/blox"

// Network creates a Network Lucide icon.
func Network(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-network", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("6"), x.X("16"), x.Y("16"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("6"), x.X("2"), x.Y("16"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("6"), x.X("9"), x.Y("2"), x.Rx("1"))),
		x.Child(x.Path(x.D("M5 16v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3"))),
		x.Child(x.Path(x.D("M12 12V8"))),
	)
	return x.Svg(svgArgs...)
}
