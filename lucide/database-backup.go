package lucide

import x "github.com/plainkit/blox"

// DatabaseBackup creates a Database Backup Lucide icon.
func DatabaseBackup(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-database-backup", args)
	svgArgs = append(svgArgs,
		x.Child(x.Ellipse(x.EllipseCx("12"), x.EllipseCy("5"), x.EllipseRx("9"), x.EllipseRy("3"))),
		x.Child(x.Path(x.D("M3 12a9 3 0 0 0 5 2.69"))),
		x.Child(x.Path(x.D("M21 9.3V5"))),
		x.Child(x.Path(x.D("M3 5v14a9 3 0 0 0 6.47 2.88"))),
		x.Child(x.Path(x.D("M12 12v4h4"))),
		x.Child(x.Path(x.D("M13 20a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16"))),
	)
	return x.Svg(svgArgs...)
}
