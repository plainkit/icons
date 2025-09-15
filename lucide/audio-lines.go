package lucide

import x "github.com/plainkit/html"

// AudioLines creates a Audio Lines Lucide icon.
func AudioLines(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-audio-lines", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 10v3"))),
		x.Child(x.Path(x.D("M6 6v11"))),
		x.Child(x.Path(x.D("M10 3v18"))),
		x.Child(x.Path(x.D("M14 8v7"))),
		x.Child(x.Path(x.D("M18 5v13"))),
		x.Child(x.Path(x.D("M22 10v3"))),
	)
	return x.Svg(svgArgs...)
}
