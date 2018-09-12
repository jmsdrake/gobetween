	/**
 * weight.go - weight balance impl
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

package balance

import (
	"../core"
	"errors"
	"math/rand"
	"../logging"
)

/**
 * Weight balancer
 */
type WeightBalancer struct{}

/**
 * Elect backend based on weight strategy
 */
func (b *WeightBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {

	if len(backends) == 0 {
		return nil, errors.New("Can't elect backend, Backends empty")
	}

	totalWeight := 0
	log := logging.For("elect/weight")

	for _, backend := range backends {
		if backend.Weight < 0 {
			return nil, errors.New("Invalid backend weight <0")
		}
		totalWeight += backend.Weight
	}

	if totalWeight == 0 {
		if len(backends) == 1 {
			for _, backend := range backends {
				return backend, nil
			}
		} else {
			r1 := rand.Intn(len(backends)-1)
			pos := 0
		
			for _, backend := range backends {

				if r1 >= pos {
					return backend, nil
				}
				pos += 1

			}

		}

	} else {
		log.Warn("total weight is:", totalWeight)

		r2 := rand.Intn(totalWeight)
		pos := 0

		for _, backend := range backends {
			pos += backend.Weight
			if r2 >= pos {
				continue
			}
			return backend, nil
		}
	}
	return nil, errors.New("Can't elect backend")
}
