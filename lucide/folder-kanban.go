package lucide

import x "github.com/plainkit/html"

// FolderKanban creates a Folder Kanban Lucide icon.
func FolderKanban(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-kanban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"))),
		x.Child(x.Path(x.D("M8 10v4"))),
		x.Child(x.Path(x.D("M12 10v2"))),
		x.Child(x.Path(x.D("M16 10v6"))),
	)
	return x.Svg(svgArgs...)
}
