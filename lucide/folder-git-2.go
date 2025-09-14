package lucide

import x "github.com/bloxui/blox"

// FolderGit2 creates a Folder Git 2 Lucide icon.
func FolderGit2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-git-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v5"))),
		x.Child(x.Circle(x.Cx("13"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("M18 19c-2.8 0-5-2.2-5-5v8"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
