package lucide

import x "github.com/bloxui/blox"

// HardHat creates a Hard Hat Lucide icon.
func HardHat(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hard-hat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10V5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5"))),
		x.Child(x.Path(x.D("M14 6a6 6 0 0 1 6 6v3"))),
		x.Child(x.Path(x.D("M4 15v-3a6 6 0 0 1 6-6"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("4"), x.X("2"), x.Y("15"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
