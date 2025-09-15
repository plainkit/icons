package lucide

import x "github.com/plainkit/html"

// FolderPen creates a Folder Pen Lucide icon.
func FolderPen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-pen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 11.5V5a2 2 0 0 1 2-2h3.9c.7 0 1.3.3 1.7.9l.8 1.2c.4.6 1 .9 1.7.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-9.5"))),
		x.Child(x.Path(x.D("M11.378 13.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
	)
	return x.Svg(svgArgs...)
}
