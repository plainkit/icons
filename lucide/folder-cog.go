package lucide

import x "github.com/bloxui/blox"

// FolderCog creates a Folder Cog Lucide icon.
func FolderCog(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-folder-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.3 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.98a2 2 0 0 1 1.69.9l.66 1.2A2 2 0 0 0 12 6h8a2 2 0 0 1 2 2v3.3"))),
		x.Child(x.Path(x.D("m14.305 19.53.923-.382"))),
		x.Child(x.Path(x.D("m15.228 16.852-.923-.383"))),
		x.Child(x.Path(x.D("m16.852 15.228-.383-.923"))),
		x.Child(x.Path(x.D("m16.852 20.772-.383.924"))),
		x.Child(x.Path(x.D("m19.148 15.228.383-.923"))),
		x.Child(x.Path(x.D("m19.53 21.696-.382-.924"))),
		x.Child(x.Path(x.D("m20.772 16.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 19.148.924.383"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
