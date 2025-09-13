package lucide

import x "github.com/bloxui/blox"

// BringToFront creates a Bring To Front Lucide icon.
func BringToFront(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bring-to-front", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.X("8"), x.Y("8"), x.RectWidth("8"), x.RectHeight("8"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M14 20a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2"))),
	)
	return x.Svg(svgArgs...)
}
