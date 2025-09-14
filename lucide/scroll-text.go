package lucide

import x "github.com/bloxui/blox"

// ScrollText creates a Scroll Text Lucide icon.
func ScrollText(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scroll-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 12h-5"))),
		x.Child(x.Path(x.D("M15 8h-5"))),
		x.Child(x.Path(x.D("M19 17V5a2 2 0 0 0-2-2H4"))),
		x.Child(x.Path(x.D("M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"))),
	)
	return x.Svg(svgArgs...)
}
