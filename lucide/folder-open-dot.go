package lucide

import x "github.com/bloxui/blox"

// FolderOpenDot creates a Folder Open Dot Lucide icon.
func FolderOpenDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-open-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 14 1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"))),
		x.Child(x.Circle(x.Cx("14"), x.Cy("15"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
