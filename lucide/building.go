package lucide

import x "github.com/bloxui/blox"

// Building creates a Building Lucide icon.
func Building(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-building", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 10h.01"))),
		x.Child(x.Path(x.D("M12 14h.01"))),
		x.Child(x.Path(x.D("M12 6h.01"))),
		x.Child(x.Path(x.D("M16 10h.01"))),
		x.Child(x.Path(x.D("M16 14h.01"))),
		x.Child(x.Path(x.D("M16 6h.01"))),
		x.Child(x.Path(x.D("M8 10h.01"))),
		x.Child(x.Path(x.D("M8 14h.01"))),
		x.Child(x.Path(x.D("M8 6h.01"))),
		x.Child(x.Path(x.D("M9 22v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
