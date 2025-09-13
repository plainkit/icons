package lucide

import x "github.com/bloxui/blox"

// FileMusic creates a File Music Lucide icon.
func FileMusic(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-music", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v8.4"))),
		x.Child(x.Path(x.D("M8 18v-7.7L16 9v7"))),
		x.Child(x.Circle(x.Cx("14"), x.Cy("16"), x.R("2"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("18"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
