package lucide

import x "github.com/plainkit/html"

// FolderClock creates a Folder Clock Lucide icon.
func FolderClock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 14v2.2l1.6 1"))),
		x.Child(x.Path(x.D("M7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("16"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
