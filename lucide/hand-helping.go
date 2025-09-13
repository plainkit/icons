package lucide

import x "github.com/bloxui/blox"

// HandHelping creates a Hand Helping Lucide icon.
func HandHelping(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hand-helping", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 12h2a2 2 0 1 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 14"))),
		x.Child(x.Path(x.D("m7 18 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a2 2 0 0 0-2.75-2.91l-4.2 3.9"))),
		x.Child(x.Path(x.D("m2 13 6 6"))),
	)
	return x.Svg(svgArgs...)
}
