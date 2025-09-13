package lucide

import x "github.com/bloxui/blox"

// Popsicle creates a Popsicle Lucide icon.
func Popsicle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-popsicle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1Z"))),
		x.Child(x.Path(x.D("m22 22-5.5-5.5"))),
	)
	return x.Svg(svgArgs...)
}
