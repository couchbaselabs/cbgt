//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package main

// A planner assigns partitions to cbft's and to PIndexes on each cbft.
func (mgr *Manager) PlannerLoop() {
	for _ = range mgr.plannerCh {
		// TODO: implement the grand plans for the planner.
		// First gen planner should keep it simple, such as...
		// - a single Feed for every datasource node.
		// - a Feed might "fan out" to multiple Streams/PIndexes.
		// - have a single PIndex for all datasource partitions
		//   (vbuckets) to start.
	}
}
