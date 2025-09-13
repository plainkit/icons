package lucide

import x "github.com/bloxui/blox"

// FileJson2 creates a File Json 2 Lucide icon.
func FileJson2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-json-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M4 12a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1 1 1 0 0 1 1 1v1a1 1 0 0 0 1 1"))),
		x.Child(x.Path(x.D("M8 18a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1 1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"))),
	)
	return x.Svg(svgArgs...)
}
