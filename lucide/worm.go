package lucide

import x "github.com/plainkit/html"

// Worm creates a Worm Lucide icon.
func Worm(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-worm", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 12-1.5 3"))),
		x.Child(x.Path(x.D("M19.63 18.81 22 20"))),
		x.Child(x.Path(x.D("M6.47 8.23a1.68 1.68 0 0 1 2.44 1.93l-.64 2.08a6.76 6.76 0 0 0 10.16 7.67l.42-.27a1 1 0 1 0-2.73-4.21l-.42.27a1.76 1.76 0 0 1-2.63-1.99l.64-2.08A6.66 6.66 0 0 0 3.94 3.9l-.7.4a1 1 0 1 0 2.55 4.34z"))),
	)
	return x.Svg(svgArgs...)
}
