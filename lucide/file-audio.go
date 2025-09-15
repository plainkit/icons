package lucide

import x "github.com/plainkit/blox"

// FileAudio creates a File Audio Lucide icon.
func FileAudio(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-audio", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.5 22h.5a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M2 19a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0v-4a6 6 0 0 1 12 0v4a2 2 0 1 1-4 0v-1a2 2 0 1 1 4 0"))),
	)
	return x.Svg(svgArgs...)
}
