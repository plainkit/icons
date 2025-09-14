package lucide

import x "github.com/bloxui/blox"

// FileAudio2 creates a File Audio 2 Lucide icon.
func FileAudio2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-audio-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v2"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Circle(x.Cx("3"), x.Cy("17"), x.R("1"))),
		x.Child(x.Path(x.D("M2 17v-3a4 4 0 0 1 8 0v3"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("17"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
