package lucide

import x "github.com/bloxui/blox"

// ListStart creates a List Start Lucide icon.
func ListStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h6"))),
		x.Child(x.Path(x.D("M3 12h13"))),
		x.Child(x.Path(x.D("M3 19h13"))),
		x.Child(x.Path(x.D("m16 8-3-3 3-3"))),
		x.Child(x.Path(x.D("M21 19V7a2 2 0 0 0-2-2h-6"))),
	)
	return x.Svg(svgArgs...)
}
