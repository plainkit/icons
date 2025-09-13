package lucide

import x "github.com/bloxui/blox"

// FileLock2 creates a File Lock 2 Lucide icon.
func FileLock2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-lock-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v1"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Rect(x.X("2"), x.Y("13"), x.RectWidth("8"), x.RectHeight("5"), x.Rx("1"))),
		x.Child(x.Path(x.D("M8 13v-2a2 2 0 1 0-4 0v2"))),
	)
	return x.Svg(svgArgs...)
}
