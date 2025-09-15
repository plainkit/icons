package lucide

import x "github.com/plainkit/html"

// FolderSearch2 creates a Folder Search 2 Lucide icon.
func FolderSearch2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-search-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11.5"), x.Cy("12.5"), x.R("2.5"))),
		x.Child(x.Path(x.D("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		x.Child(x.Path(x.D("M13.3 14.3 15 16"))),
	)
	return x.Svg(svgArgs...)
}
