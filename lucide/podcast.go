package lucide

import x "github.com/bloxui/blox"

// Podcast creates a Podcast Lucide icon.
func Podcast(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-podcast", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 17a1 1 0 1 0-2 0l.5 4.5a0.5 0.5 0 0 0 1 0z"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("M16.85 18.58a9 9 0 1 0-9.7 0"))),
		x.Child(x.Path(x.D("M8 14a5 5 0 1 1 8 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("11"), x.R("1"), x.Fill("currentColor"))),
	)
	return x.Svg(svgArgs...)
}
