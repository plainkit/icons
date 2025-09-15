package lucide

import x "github.com/plainkit/html"

// CookingPot creates a Cooking Pot Lucide icon.
func CookingPot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cooking-pot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("M20 12v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"))),
		x.Child(x.Path(x.D("m4 8 16-4"))),
		x.Child(x.Path(x.D("m8.86 6.78-.45-1.81a2 2 0 0 1 1.45-2.43l1.94-.48a2 2 0 0 1 2.43 1.46l.45 1.8"))),
	)
	return x.Svg(svgArgs...)
}
