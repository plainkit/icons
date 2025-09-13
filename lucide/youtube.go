package lucide

import x "github.com/bloxui/blox"

// Youtube creates a Youtube Lucide icon.
func Youtube(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-youtube", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.5 17a24.12 24.12 0 0 1 0-10 2 2 0 0 1 1.4-1.4 49.56 49.56 0 0 1 16.2 0A2 2 0 0 1 21.5 7a24.12 24.12 0 0 1 0 10 2 2 0 0 1-1.4 1.4 49.55 49.55 0 0 1-16.2 0A2 2 0 0 1 2.5 17"))),
		x.Child(x.Path(x.D("m10 15 5-3-5-3z"))),
	)
	return x.Svg(svgArgs...)
}
