package resources

import (
	gossResource "github.com/malnick/goss/resource"
)

var (
	s1 = Service{
		Service: gossResource.Service{
			Title:   "service1",
			Running: true,
		},
	}

	s2 = Service{
		Service: gossResource.Service{
			Title:   "service2",
			Running: true,
		},
	}

	s3 = Service{
		Service: gossResource.Service{
			Title:   "service3",
			Running: true,
		},
	}
)
