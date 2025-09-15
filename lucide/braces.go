package lucide

import x "github.com/plainkit/blox"

// Braces creates a Braces Lucide icon.
func Braces(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-braces", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2 2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1"))),
		x.Child(x.Path(x.D("M16 21h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"))),
	)
	return x.Svg(svgArgs...)
}
