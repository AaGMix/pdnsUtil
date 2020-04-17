package pdnsutil

import "github.com/AaGMix/pdnsUtil/zones"

type Control interface {
	Zones() zones.Control
}
