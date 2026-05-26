/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// modified from https://knative.dev/pkg/tree/master/tracker
// The version provided by knative/pkg forces tracking on namespace scoped
// object an in our case the ClusterBuilder is a cluster scoped
// object that need to be tracked

package tracker

import (
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"github.com/pivotal/kpack/pkg/reconciler"
)

func New(callback func(types.NamespacedName), lease time.Duration) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

type Tracker struct {
	m sync.Mutex
	// objects maps from an object reference to the set of
	// keys for objects watching it.
	objects map[string]set

	// kinds maps from group version kind to the set of
	// keys for objects watching it.
	kinds map[string]set

	// The amount of time that an object may watch another
	// before having to renew the lease.
	leaseDuration time.Duration

	cb func(types.NamespacedName)
}

// set is a map from keys to expirations
type set map[types.NamespacedName]time.Time

// Track implements Interface.
func (i *Tracker) Track(ref reconciler.Key, obj types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

// Overwrite the key with a new expiration.

func isExpired(expiry time.Time) bool { _ = "STUB: not implemented"; return false }

func (i *Tracker) TrackKind(kind schema.GroupKind, obj types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

// Overwrite the key with a new expiration.

// OnChanged implements Interface.
func (i *Tracker) OnChanged(obj interface{}) { _ = "STUB: not implemented"; return }

// TODO(mattmoor): Consider locking the mapping (global) for a
// smaller scope and leveraging a per-set lock to guard its access.

func (i *Tracker) notify(mapping map[string]set, key string) { _ = "STUB: not implemented"; return }

// TODO(mattmoor): We should consider logging here.

// If the expiration has lapsed, then delete the key.
