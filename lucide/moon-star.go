package lucide

import x "github.com/plainkit/blox"

// MoonStar creates a Moon Star Lucide icon.
func MoonStar(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-moon-star", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 5h4"))),
		x.Child(x.Path(x.D("M20 3v4"))),
		x.Child(x.Path(x.D("M20.985 12.486a9 9 0 1 1-9.473-9.472c.405-.022.617.46.402.803a6 6 0 0 0 8.268 8.268c.344-.215.825-.004.803.401"))),
	)
	return x.Svg(svgArgs...)
}
