package lucide

import x "github.com/bloxui/blox"

// HandHeart creates a Hand Heart Lucide icon.
func HandHeart(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hand-heart", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 14h2a2 2 0 0 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 16"))),
		x.Child(x.Path(x.D("m14.45 13.39 5.05-4.694C20.196 8 21 6.85 21 5.75a2.75 2.75 0 0 0-4.797-1.837.276.276 0 0 1-.406 0A2.75 2.75 0 0 0 11 5.75c0 1.2.802 2.248 1.5 2.946L16 11.95"))),
		x.Child(x.Path(x.D("m2 15 6 6"))),
		x.Child(x.Path(x.D("m7 20 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a1 1 0 0 0-2.75-2.91"))),
	)
	return x.Svg(svgArgs...)
}
