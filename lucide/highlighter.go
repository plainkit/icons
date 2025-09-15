package lucide

import x "github.com/plainkit/html"

// Highlighter creates a Highlighter Lucide icon.
func Highlighter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-highlighter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 11-6 6v3h9l3-3"))),
		x.Child(x.Path(x.D("m22 12-4.6 4.6a2 2 0 0 1-2.8 0l-5.2-5.2a2 2 0 0 1 0-2.8L14 4"))),
	)
	return x.Svg(svgArgs...)
}
