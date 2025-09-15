package lucide

import x "github.com/plainkit/blox"

// Facebook creates a Facebook Lucide icon.
func Facebook(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-facebook", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"))),
	)
	return x.Svg(svgArgs...)
}
