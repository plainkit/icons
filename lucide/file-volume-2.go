package lucide

import x "github.com/bloxui/blox"

// FileVolume2 creates a File Volume 2 Lucide icon.
func FileVolume2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-volume-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M8 15h.01"))),
		x.Child(x.Path(x.D("M11.5 13.5a2.5 2.5 0 0 1 0 3"))),
		x.Child(x.Path(x.D("M15 12a5 5 0 0 1 0 6"))),
	)
	return x.Svg(svgArgs...)
}
