package resources

import (
	gossResource "github.com/malnick/goss/resource"
)

var (
	p1 = Package{
		Package: gossResource.Package{
			Title:     "pkg1",
			Installed: true,
		},
	}

	p2 = Package{
		Package: gossResource.Package{
			Title:     "pkg2",
			Installed: true,
		},
	}

	p3 = Package{
		Package: gossResource.Package{
			Title:     "pkg3",
			Installed: true,
		},
	}
)
