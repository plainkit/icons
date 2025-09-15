package lucide

import x "github.com/plainkit/html"

// LightbulbOff creates a Lightbulb Off Lucide icon.
func LightbulbOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lightbulb-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16.8 11.2c.8-.9 1.2-2 1.2-3.2a6 6 0 0 0-9.3-5"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M6.3 6.3a4.67 4.67 0 0 0 1.2 5.2c.7.7 1.3 1.5 1.5 2.5"))),
		x.Child(x.Path(x.D("M9 18h6"))),
		x.Child(x.Path(x.D("M10 22h4"))),
	)
	return x.Svg(svgArgs...)
}
