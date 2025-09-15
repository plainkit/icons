package lucide

import x "github.com/plainkit/html"

// FolderGit creates a Folder Git Lucide icon.
func FolderGit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-git", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("2"))),
		x.Child(x.Path(x.D("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		x.Child(x.Path(x.D("M14 13h3"))),
		x.Child(x.Path(x.D("M7 13h3"))),
	)
	return x.Svg(svgArgs...)
}
