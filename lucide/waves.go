package lucide

import x "github.com/plainkit/html"

// Waves creates a Waves Lucide icon.
func Waves(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-waves", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 6c.6.5 1.2 1 2.5 1C7 7 7 5 9.5 5c2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
		x.Child(x.Path(x.D("M2 12c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
		x.Child(x.Path(x.D("M2 18c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
	)
	return x.Svg(svgArgs...)
}
