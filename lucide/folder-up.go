package lucide

import x "github.com/plainkit/html"

// FolderUp creates a Folder Up Lucide icon.
func FolderUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		x.Child(x.Path(x.D("M12 10v6"))),
		x.Child(x.Path(x.D("m9 13 3-3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
