package main

func renewCerts(swarmdir, unlockKey string) error {
	_ = "STUB: not implemented"
	// First, load the existing cert.  We don't actually bother to check if
	// it's expired - this will just obtain a new cert anyway.
	return nil
}

// We need to make sure when renewing that we provide the same CN (node ID),
// OU (role), and org (swarm cluster ID) when getting a new certificate

// Load up the raft data on disk

// If there's a snapshot, get the cluster from it

// It's possible there's no snapshot yet, or the cluster has been updated
// since the last snapshot, so also read from the WALs

// There should always be a cluster and CA cert, unless the raft store has been
// catastrophcially corrupted, but it's possible that there is no CA key because
// the cluster used an external CA.

// Issue a new certificate that expires at the configured expiry time.
