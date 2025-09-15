package lucide

import x "github.com/plainkit/html"

// FolderSearch creates a Folder Search Lucide icon.
func FolderSearch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v4.1"))),
		x.Child(x.Path(x.D("m21 21-1.9-1.9"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("17"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
