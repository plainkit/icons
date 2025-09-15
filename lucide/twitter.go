package lucide

import x "github.com/plainkit/blox"

// Twitter creates a Twitter Lucide icon.
func Twitter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-twitter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z"))),
	)
	return x.Svg(svgArgs...)
}
