package lucide

import x "github.com/bloxui/blox"

// FileBadge creates a File Badge Lucide icon.
func FileBadge(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-badge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22h6a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3.072"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m6.69 16.479 1.29 4.88a.5.5 0 0 1-.698.591l-1.843-.849a1 1 0 0 0-.88.001l-1.846.85a.5.5 0 0 1-.693-.593l1.29-4.88"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("14"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
