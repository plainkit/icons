package lucide

import x "github.com/bloxui/blox"

// Scroll creates a Scroll Lucide icon.
func Scroll(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-scroll", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 17V5a2 2 0 0 0-2-2H4"))),
		x.Child(x.Path(x.D("M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"))),
	)
	return x.Svg(svgArgs...)
}
