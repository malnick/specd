package resources

import (
	gossResource "github.com/malnick/goss/resource"
)

var (
	f1 = File{
		File: gossResource.File{
			Title:  "test1",
			Mode:   "0600",
			Exists: true,
		},
	}

	f2 = File{
		File: gossResource.File{
			Title:  "test2",
			Exists: true,
		},
	}

	f3 = File{
		File: gossResource.File{
			Title:  "test3",
			Exists: true,
		},
	}
)
