package lucide

import x "github.com/plainkit/blox"

// BookAlert creates a Book Alert Lucide icon.
func BookAlert(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13h.01"))),
		x.Child(x.Path(x.D("M12 6v3"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
	)
	return x.Svg(svgArgs...)
}
