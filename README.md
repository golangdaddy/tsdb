# metrick

Multi-dimensional metrics for Kubernetes clusters.

This package consists of a client, and a server, which can be used to collect
and analyse metrics created by golang applications.

## The Ingester

The Ingester is a server that caches events in memory for fast querying.

It uses merkle trees and bloom filters to speed up querying, as to not use a
full graph structure to query, which is beneficial for the garbage collector.

It needs to scale vertically to handle around 20k events per second.

The following code will initialise the package to ingest metrics found in the VolumePath folder.

```
package main

import (
	"github.com/DianomiLtd/metrick"
)

func main() {

	tree := metrick.NewTree(
		metrick.TreeSettings{
			VolumePath: "/my-nfs-mount",
			BufferDuration: 24 * 7 * time.Hour,
		},
	)

	<- make(chan bool)

}
```

## Writing Logs

Writing logs is done in the go applications using the following code.


```
	client.Publish(
		metrick.NewMetric(

		),
	)

```
