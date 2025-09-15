package lucide

import x "github.com/plainkit/blox"

// Receipt creates a Receipt Lucide icon.
func Receipt(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"))),
		x.Child(x.Path(x.D("M12 17.5v-11"))),
	)
	return x.Svg(svgArgs...)
}
