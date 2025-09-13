package lucide

import x "github.com/bloxui/blox"

// Piano creates a Piano Lucide icon.
func Piano(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-piano", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.5 8c-1.4 0-2.6-.8-3.2-2A6.87 6.87 0 0 0 2 9v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8.5C22 9.6 20.4 8 18.5 8"))),
		x.Child(x.Path(x.D("M2 14h20"))),
		x.Child(x.Path(x.D("M6 14v4"))),
		x.Child(x.Path(x.D("M10 14v4"))),
		x.Child(x.Path(x.D("M14 14v4"))),
		x.Child(x.Path(x.D("M18 14v4"))),
	)
	return x.Svg(svgArgs...)
}
