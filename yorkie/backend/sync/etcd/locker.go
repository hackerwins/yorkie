/*
 * Copyright 2021 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package etcd

import (
	"context"

	"go.etcd.io/etcd/clientv3/concurrency"

	"github.com/yorkie-team/yorkie/pkg/log"
)

type internalLocker struct {
	session *concurrency.Session
	mu      *concurrency.Mutex
}

// Lock locks a mutex.
func (il *internalLocker) Lock(ctx context.Context) error {
	if err := il.mu.Lock(ctx); err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

// Unlock unlocks the mutex.
func (il *internalLocker) Unlock(ctx context.Context) error {
	if err := il.mu.Unlock(ctx); err != nil {
		log.Logger.Error(err)
		return err
	}
	if err := il.session.Close(); err != nil {
		return err
	}

	return nil
}
