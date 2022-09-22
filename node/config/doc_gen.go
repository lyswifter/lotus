// Code generated by github.com/filecoin-project/lotus/node/config/cfgdocgen. DO NOT EDIT.

package config

type DocField struct {
	Num     string
	Type    string
	Comment string
}

var Doc = map[string][]DocField{
	"API": []DocField{
		{
			Num:  "ListenAddress",
			Type: "string",

			Comment: `Binding address for the Lotus API`,
		},
		{
			Num:  "RemoteListenAddress",
			Type: "string",

			Comment: ``,
		},
		{
			Num:  "Timeout",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Backup": []DocField{
		{
			Num:  "DisableMetadataLog",
			Type: "bool",

			Comment: `When set to true disables metadata log (.lotus/kvlog). This can save disk
space by reducing metadata redundancy.

Note that in case of metadata corruption it might be much harder to recover
your node if metadata log is disabled`,
		},
	},
	"BatchFeeConfig": []DocField{
		{
			Num:  "Base",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Num:  "PerSector",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"Chainstore": []DocField{
		{
			Num:  "EnableSplitstore",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "Splitstore",
			Type: "Splitstore",

			Comment: ``,
		},
	},
	"Client": []DocField{
		{
			Num:  "UseIpfs",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "IpfsOnlineMode",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "IpfsMAddr",
			Type: "string",

			Comment: ``,
		},
		{
			Num:  "IpfsUseForRetrieval",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for storage deals`,
		},
		{
			Num:  "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for retrieval deals`,
		},
		{
			Num:  "OffChainRetrieval",
			Type: "bool",

			Comment: `Require that retrievals perform no on-chain operations. Paid retrievals
without existing payment channels with available funds will fail instead
of automatically performing on-chain operations.`,
		},
	},
	"ClusterRaftConfig": []DocField{
		{
			Num:  "ClusterModeEnabled",
			Type: "bool",

			Comment: `config to enabled node cluster with raft consensus`,
		},
		{
			Num:  "HostShutdown",
			Type: "bool",

			Comment: `will shutdown libp2p host on shutdown. Useful for testing`,
		},
		{
			Num:  "DataFolder",
			Type: "string",

			Comment: `A folder to store Raft's data.`,
		},
		{
			Num:  "InitPeerset",
			Type: "[]peer.ID",

			Comment: `InitPeerset provides the list of initial cluster peers for new Raft
peers (with no prior state). It is ignored when Raft was already
initialized or when starting in staging mode.`,
		},
		{
			Num:  "WaitForLeaderTimeout",
			Type: "Duration",

			Comment: `LeaderTimeout specifies how long to wait for a leader before
failing an operation.`,
		},
		{
			Num:  "NetworkTimeout",
			Type: "Duration",

			Comment: `NetworkTimeout specifies how long before a Raft network
operation is timed out`,
		},
		{
			Num:  "CommitRetries",
			Type: "int",

			Comment: `CommitRetries specifies how many times we retry a failed commit until
we give up.`,
		},
		{
			Num:  "CommitRetryDelay",
			Type: "Duration",

			Comment: `How long to wait between retries`,
		},
		{
			Num:  "BackupsRotate",
			Type: "int",

			Comment: `BackupsRotate specifies the maximum number of Raft's DataFolder
copies that we keep as backups (renaming) after cleanup.`,
		},
		{
			Num:  "DatastoreNamespace",
			Type: "string",

			Comment: `Namespace to use when writing keys to the datastore`,
		},
		{
			Num:  "RaftConfig",
			Type: "*hraft.Config",

			Comment: `A Hashicorp Raft's configuration object.`,
		},
		{
			Num:  "Tracing",
			Type: "bool",

			Comment: `Tracing enables propagation of contexts across binary boundaries.`,
		},
	},
	"Common": []DocField{
		{
			Num:  "API",
			Type: "API",

			Comment: ``,
		},
		{
			Num:  "Backup",
			Type: "Backup",

			Comment: ``,
		},
		{
			Num:  "Logging",
			Type: "Logging",

			Comment: ``,
		},
		{
			Num:  "Libp2p",
			Type: "Libp2p",

			Comment: ``,
		},
		{
			Num:  "Pubsub",
			Type: "Pubsub",

			Comment: ``,
		},
	},
	"DAGStoreConfig": []DocField{
		{
			Num:  "RootDir",
			Type: "string",

			Comment: `Path to the dagstore root directory. This directory contains three
subdirectories, which can be symlinked to alternative locations if
need be:
- ./transients: caches unsealed deals that have been fetched from the
storage subsystem for serving retrievals.
- ./indices: stores shard indices.
- ./datastore: holds the KV store tracking the state of every shard
known to the DAG store.
Default value: <LOTUS_MARKETS_PATH>/dagstore (split deployment) or
<LOTUS_MINER_PATH>/dagstore (monolith deployment)`,
		},
		{
			Num:  "MaxConcurrentIndex",
			Type: "int",

			Comment: `The maximum amount of indexing jobs that can run simultaneously.
0 means unlimited.
Default value: 5.`,
		},
		{
			Num:  "MaxConcurrentReadyFetches",
			Type: "int",

			Comment: `The maximum amount of unsealed deals that can be fetched simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Num:  "MaxConcurrentUnseals",
			Type: "int",

			Comment: `The maximum amount of unseals that can be processed simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Num:  "MaxConcurrencyStorageCalls",
			Type: "int",

			Comment: `The maximum number of simultaneous inflight API calls to the storage
subsystem.
Default value: 100.`,
		},
		{
			Num:  "GCInterval",
			Type: "Duration",

			Comment: `The time between calls to periodic dagstore GC, in time.Duration string
representation, e.g. 1m, 5m, 1h.
Default value: 1 minute.`,
		},
	},
	"DealmakingConfig": []DocField{
		{
			Num:  "ConsiderOnlineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept online deals`,
		},
		{
			Num:  "ConsiderOfflineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline deals`,
		},
		{
			Num:  "ConsiderOnlineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept retrieval deals`,
		},
		{
			Num:  "ConsiderOfflineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline retrieval deals`,
		},
		{
			Num:  "ConsiderVerifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept verified deals`,
		},
		{
			Num:  "ConsiderUnverifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept unverified deals`,
		},
		{
			Num:  "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Num:  "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Num:  "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Num:  "PublishMsgPeriod",
			Type: "Duration",

			Comment: `When a deal is ready to publish, the amount of time to wait for more
deals to be ready to publish before publishing them all as a batch`,
		},
		{
			Num:  "MaxDealsPerPublishMsg",
			Type: "uint64",

			Comment: `The maximum number of deals to include in a single PublishStorageDeals
message`,
		},
		{
			Num:  "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Num:  "MaxStagingDealsBytes",
			Type: "int64",

			Comment: `The maximum allowed disk usage size in bytes of staging deals not yet
passed to the sealing node by the markets service. 0 is unlimited.`,
		},
		{
			Num:  "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for storage deals`,
		},
		{
			Num:  "SimultaneousTransfersForStoragePerClient",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers from any single client
for storage deals.
Unset by default (0), and values higher than SimultaneousTransfersForStorage
will have no effect; i.e. the total number of simultaneous data transfers
across all storage clients is bound by SimultaneousTransfersForStorage
regardless of this number.`,
		},
		{
			Num:  "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for retrieval deals`,
		},
		{
			Num:  "StartEpochSealingBuffer",
			Type: "uint64",

			Comment: `Minimum start epoch buffer to give time for sealing of sector with deal.`,
		},
		{
			Num:  "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Num:  "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Num:  "RetrievalPricing",
			Type: "*RetrievalPricing",

			Comment: ``,
		},
	},
	"FeeConfig": []DocField{
		{
			Num:  "DefaultMaxFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"FullNode": []DocField{
		{
			Num:  "Client",
			Type: "Client",

			Comment: ``,
		},
		{
			Num:  "Wallet",
			Type: "Wallet",

			Comment: ``,
		},
		{
			Num:  "Fees",
			Type: "FeeConfig",

			Comment: ``,
		},
		{
			Num:  "Chainstore",
			Type: "Chainstore",

			Comment: ``,
		},
		{
			Num:  "Raft",
			Type: "ClusterRaftConfig",

			Comment: ``,
		},
	},
	"IndexProviderConfig": []DocField{
		{
			Num:  "Enable",
			Type: "bool",

			Comment: `Enable set whether to enable indexing announcement to the network and expose endpoints that
allow indexer nodes to process announcements. Enabled by default.`,
		},
		{
			Num:  "EntriesCacheCapacity",
			Type: "int",

			Comment: `EntriesCacheCapacity sets the maximum capacity to use for caching the indexing advertisement
entries. Defaults to 1024 if not specified. The cache is evicted using LRU policy. The
maximum storage used by the cache is a factor of EntriesCacheCapacity, EntriesChunkSize and
the length of multihashes being advertised. For example, advertising 128-bit long multihashes
with the default EntriesCacheCapacity, and EntriesChunkSize means the cache size can grow to
256MiB when full.`,
		},
		{
			Num:  "EntriesChunkSize",
			Type: "int",

			Comment: `EntriesChunkSize sets the maximum number of multihashes to include in a single entries chunk.
Defaults to 16384 if not specified. Note that chunks are chained together for indexing
advertisements that include more multihashes than the configured EntriesChunkSize.`,
		},
		{
			Num:  "TopicName",
			Type: "string",

			Comment: `TopicName sets the topic name on which the changes to the advertised content are announced.
If not explicitly specified, the topic name is automatically inferred from the network name
in following format: '/indexer/ingest/<network-name>'
Defaults to empty, which implies the topic name is inferred from network name.`,
		},
		{
			Num:  "PurgeCacheOnStart",
			Type: "bool",

			Comment: `PurgeCacheOnStart sets whether to clear any cached entries chunks when the provider engine
starts. By default, the cache is rehydrated from previously cached entries stored in
datastore if any is present.`,
		},
	},
	"Libp2p": []DocField{
		{
			Num:  "ListenAddresses",
			Type: "[]string",

			Comment: `Binding address for the libp2p host - 0 means random port.
Format: multiaddress; see https://multiformats.io/multiaddr/`,
		},
		{
			Num:  "AnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to explicitally announce to other peers. If not specified,
all interface addresses are announced
Format: multiaddress`,
		},
		{
			Num:  "NoAnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to not announce
Format: multiaddress`,
		},
		{
			Num:  "BootstrapPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Num:  "ProtectedPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Num:  "DisableNatPortMap",
			Type: "bool",

			Comment: `When not disabled (default), lotus asks NAT devices (e.g., routers), to
open up an external port and forward it to the port lotus is running on.
When this works (i.e., when your router supports NAT port forwarding),
it makes the local lotus node accessible from the public internet`,
		},
		{
			Num:  "ConnMgrLow",
			Type: "uint",

			Comment: `ConnMgrLow is the number of connections that the basic connection manager
will trim down to.`,
		},
		{
			Num:  "ConnMgrHigh",
			Type: "uint",

			Comment: `ConnMgrHigh is the number of connections that, when exceeded, will trigger
a connection GC operation. Note: protected/recently formed connections don't
count towards this limit.`,
		},
		{
			Num:  "ConnMgrGrace",
			Type: "Duration",

			Comment: `ConnMgrGrace is a time duration that new connections are immune from being
closed by the connection manager.`,
		},
	},
	"Logging": []DocField{
		{
			Num:  "SubsystemLevels",
			Type: "map[string]string",

			Comment: `SubsystemLevels specify per-subsystem log levels`,
		},
	},
	"MinerAddressConfig": []DocField{
		{
			Num:  "PreCommitControl",
			Type: "[]string",

			Comment: `Addresses to send PreCommit messages from`,
		},
		{
			Num:  "CommitControl",
			Type: "[]string",

			Comment: `Addresses to send Commit messages from`,
		},
		{
			Num:  "TerminateControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Num:  "DealPublishControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Num:  "DisableOwnerFallback",
			Type: "bool",

			Comment: `DisableOwnerFallback disables usage of the owner address for messages
sent automatically`,
		},
		{
			Num:  "DisableWorkerFallback",
			Type: "bool",

			Comment: `DisableWorkerFallback disables usage of the worker address for messages
sent automatically, if control addresses are configured.
A control address that doesn't have enough funds will still be chosen
over the worker address if this flag is set.`,
		},
	},
	"MinerFeeConfig": []DocField{
		{
			Num:  "MaxPreCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Num:  "MaxCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Num:  "MaxPreCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: `maxBatchFee = maxBase + maxPerSector * nSectors`,
		},
		{
			Num:  "MaxCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: ``,
		},
		{
			Num:  "MaxTerminateGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Num:  "MaxWindowPoStGasFee",
			Type: "types.FIL",

			Comment: `WindowPoSt is a high-value operation, so the default fee should be high.`,
		},
		{
			Num:  "MaxPublishDealsFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Num:  "MaxMarketBalanceAddFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"MinerSubsystemConfig": []DocField{
		{
			Num:  "EnableMining",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "EnableSealing",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "EnableSectorStorage",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "EnableMarkets",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "SealerApiInfo",
			Type: "string",

			Comment: ``,
		},
		{
			Num:  "SectorIndexApiInfo",
			Type: "string",

			Comment: ``,
		},
	},
	"ProvingConfig": []DocField{
		{
			Num:  "ParallelCheckLimit",
			Type: "int",

			Comment: `Maximum number of sector checks to run in parallel. (0 = unlimited)

WARNING: Setting this value too high may make the node crash by running out of stack
WARNING: Setting this value too low may make sector challenge reading much slower, resulting in failed PoSt due
to late submission.

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Num:  "DisableBuiltinWindowPoSt",
			Type: "bool",

			Comment: `Disable Window PoSt computation on the lotus-miner process even if no window PoSt workers are present.

WARNING: If no windowPoSt workers are connected, window PoSt WILL FAIL resulting in faulty sectors which will need
to be recovered. Before enabling this option, make sure your PoSt workers work correctly.

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Num:  "DisableBuiltinWinningPoSt",
			Type: "bool",

			Comment: `Disable Winning PoSt computation on the lotus-miner process even if no winning PoSt workers are present.

WARNING: If no WinningPoSt workers are connected, Winning PoSt WILL FAIL resulting in lost block rewards.
Before enabling this option, make sure your PoSt workers work correctly.`,
		},
		{
			Num:  "DisableWDPoStPreChecks",
			Type: "bool",

			Comment: `Disable WindowPoSt provable sector readability checks.

In normal operation, when preparing to compute WindowPoSt, lotus-miner will perform a round of reading challenges
from all sectors to confirm that those sectors can be proven. Challenges read in this process are discarded, as
we're only interested in checking that sector data can be read.

When using builtin proof computation (no PoSt workers, and DisableBuiltinWindowPoSt is set to false), this process
can save a lot of time and compute resources in the case that some sectors are not readable - this is caused by
the builtin logic not skipping snark computation when some sectors need to be skipped.

When using PoSt workers, this process is mostly redundant, with PoSt workers challenges will be read once, and
if challenges for some sectors aren't readable, those sectors will just get skipped.

Disabling sector pre-checks will slightly reduce IO load when proving sectors, possibly resulting in shorter
time to produce window PoSt. In setups with good IO capabilities the effect of this option on proving time should
be negligible.

NOTE: It likely is a bad idea to disable sector pre-checks in setups with no PoSt workers.

NOTE: Even when this option is enabled, recovering sectors will be checked before recovery declaration message is
sent to the chain

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Num:  "MaxPartitionsPerPoStMessage",
			Type: "int",

			Comment: `Maximum number of partitions to prove in a single SubmitWindowPoSt messace. 0 = network limit (10 in nv16)

A single partition may contain up to 2349 32GiB sectors, or 2300 64GiB sectors.

The maximum number of sectors which can be proven in a single PoSt message is 25000 in network version 16, which
means that a single message can prove at most 10 partinions

In some cases when submitting PoSt messages which are recovering sectors, the default network limit may still be
too high to fit in the block gas limit; In those cases it may be necessary to set this value to something lower
than 10; Note that setting this value lower may result in less efficient gas use - more messages will be sent,
to prove each deadline, resulting in more total gas use (but each message will have lower gas limit)

Setting this value above the network limit has no effect`,
		},
		{
			Num:  "MaxPartitionsPerRecoveryMessage",
			Type: "int",

			Comment: `In some cases when submitting DeclareFaultsRecovered messages,
there may be too many recoveries to fit in a BlockGasLimit.
In those cases it may be necessary to set this value to something low (eg 1);
Note that setting this value lower may result in less efficient gas use - more messages will be sent than needed,
resulting in more total gas use (but each message will have lower gas limit)`,
		},
	},
	"Pubsub": []DocField{
		{
			Num:  "Bootstrapper",
			Type: "bool",

			Comment: `Run the node in bootstrap-node mode`,
		},
		{
			Num:  "DirectPeers",
			Type: "[]string",

			Comment: `DirectPeers specifies peers with direct peering agreements. These peers are
connected outside of the mesh, with all (valid) message unconditionally
forwarded to them. The router will maintain open connections to these peers.
Note that the peering agreement should be reciprocal with direct peers
symmetrically configured at both ends.
Type: Array of multiaddress peerinfo strings, must include peerid (/p2p/12D3K...`,
		},
		{
			Num:  "IPColocationWhitelist",
			Type: "[]string",

			Comment: ``,
		},
		{
			Num:  "RemoteTracer",
			Type: "string",

			Comment: ``,
		},
	},
	"RetrievalPricing": []DocField{
		{
			Num:  "Strategy",
			Type: "string",

			Comment: ``,
		},
		{
			Num:  "Default",
			Type: "*RetrievalPricingDefault",

			Comment: ``,
		},
		{
			Num:  "External",
			Type: "*RetrievalPricingExternal",

			Comment: ``,
		},
	},
	"RetrievalPricingDefault": []DocField{
		{
			Num:  "VerifiedDealsFreeTransfer",
			Type: "bool",

			Comment: `VerifiedDealsFreeTransfer configures zero fees for data transfer for a retrieval deal
of a payloadCid that belongs to a verified storage deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "default".
default value is true`,
		},
	},
	"RetrievalPricingExternal": []DocField{
		{
			Num:  "Path",
			Type: "string",

			Comment: `Path of the external script that will be run to price a retrieval deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "external".`,
		},
	},
	"SealerConfig": []DocField{
		{
			Num:  "ParallelFetchLimit",
			Type: "int",

			Comment: ``,
		},
		{
			Num:  "AllowAddPiece",
			Type: "bool",

			Comment: `Local worker config`,
		},
		{
			Num:  "AllowPreCommit1",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowPreCommit2",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowCommit",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowUnseal",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowReplicaUpdate",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowProveReplicaUpdate2",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "AllowRegenSectorKey",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "LocalWorkerName",
			Type: "string",

			Comment: `LocalWorkerName specifies a custom name for the builtin worker.
If set to an empty string (default) os hostname will be used`,
		},
		{
			Num:  "Assigner",
			Type: "string",

			Comment: `Assigner specifies the worker assigner to use when scheduling tasks.
"utilization" (default) - assign tasks to workers with lowest utilization.
"spread" - assign tasks to as many distinct workers as possible.`,
		},
		{
			Num:  "DisallowRemoteFinalize",
			Type: "bool",

			Comment: `DisallowRemoteFinalize when set to true will force all Finalize tasks to
run on workers with local access to both long-term storage and the sealing
path containing the sector.
--
WARNING: Only set this if all workers have access to long-term storage
paths. If this flag is enabled, and there are workers without long-term
storage access, sectors will not be moved from them, and Finalize tasks
will appear to be stuck.
--
If you see stuck Finalize tasks after enabling this setting, check
'lotus-miner sealing sched-diag' and 'lotus-miner storage find [sector num]'`,
		},
		{
			Num:  "ResourceFiltering",
			Type: "sealer.ResourceFilteringStrategy",

			Comment: `ResourceFiltering instructs the system which resource filtering strategy
to use when evaluating tasks against this worker. An empty value defaults
to "hardware".`,
		},
	},
	"SealingConfig": []DocField{
		{
			Num:  "MaxWaitDealsSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be waiting for more deals to be packed in it before it begins sealing at any given time.
If the miner is accepting multiple deals in parallel, up to MaxWaitDealsSectors of new sectors will be created.
If more than MaxWaitDealsSectors deals are accepted in parallel, only MaxWaitDealsSectors deals will be processed in parallel
Note that setting this number too high in relation to deal ingestion rate may result in poor sector packing efficiency
0 = no limit`,
		},
		{
			Num:  "MaxSealingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new CC sectors (0 = unlimited)`,
		},
		{
			Num:  "MaxSealingSectorsForDeals",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new sectors with deals (0 = unlimited)`,
		},
		{
			Num:  "PreferNewSectorsForDeals",
			Type: "bool",

			Comment: `Prefer creating new sectors even if there are sectors Available for upgrading.
This setting combined with MaxUpgradingSectors set to a value higher than MaxSealingSectorsForDeals makes it
possible to use fast sector upgrades to handle high volumes of storage deals, while still using the simple sealing
flow when the volume of storage deals is lower.`,
		},
		{
			Num:  "MaxUpgradingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when upgrading CC sectors with deals (0 = MaxSealingSectorsForDeals)`,
		},
		{
			Num:  "CommittedCapacitySectorLifetime",
			Type: "Duration",

			Comment: `CommittedCapacitySectorLifetime is the duration a Committed Capacity (CC) sector will
live before it must be extended or converted into sector containing deals before it is
terminated. Value must be between 180-540 days inclusive`,
		},
		{
			Num:  "WaitDealsDelay",
			Type: "Duration",

			Comment: `Period of time that a newly created sector will wait for more deals to be packed in to before it starts to seal.
Sectors which are fully filled will start sealing immediately`,
		},
		{
			Num:  "AlwaysKeepUnsealedCopy",
			Type: "bool",

			Comment: `Whether to keep unsealed copies of deal data regardless of whether the client requested that. This lets the miner
avoid the relatively high cost of unsealing the data later, at the cost of more storage space`,
		},
		{
			Num:  "FinalizeEarly",
			Type: "bool",

			Comment: `Run sector finalization before submitting sector proof to the chain`,
		},
		{
			Num:  "MakeNewSectorForDeals",
			Type: "bool",

			Comment: `Whether new sectors are created to pack incoming deals
When this is set to false no new sectors will be created for sealing incoming deals
This is useful for forcing all deals to be assigned as snap deals to sectors marked for upgrade`,
		},
		{
			Num:  "MakeCCSectorsAvailable",
			Type: "bool",

			Comment: `After sealing CC sectors, make them available for upgrading with deals`,
		},
		{
			Num:  "CollateralFromMinerBalance",
			Type: "bool",

			Comment: `Whether to use available miner balance for sector collateral instead of sending it with each message`,
		},
		{
			Num:  "AvailableBalanceBuffer",
			Type: "types.FIL",

			Comment: `Minimum available balance to keep in the miner actor before sending it with messages`,
		},
		{
			Num:  "DisableCollateralFallback",
			Type: "bool",

			Comment: `Don't send collateral with messages even if there is no available balance in the miner actor`,
		},
		{
			Num:  "BatchPreCommits",
			Type: "bool",

			Comment: `enable / disable precommit batching (takes effect after nv13)`,
		},
		{
			Num:  "MaxPreCommitBatch",
			Type: "int",

			Comment: `maximum precommit batch size - batches will be sent immediately above this size`,
		},
		{
			Num:  "PreCommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Num:  "PreCommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deal in batch would start expiring`,
		},
		{
			Num:  "AggregateCommits",
			Type: "bool",

			Comment: `enable / disable commit aggregation (takes effect after nv13)`,
		},
		{
			Num:  "MinCommitBatch",
			Type: "int",

			Comment: `minimum batched commit size - batches above this size will eventually be sent on a timeout`,
		},
		{
			Num:  "MaxCommitBatch",
			Type: "int",

			Comment: `maximum batched commit size - batches will be sent immediately above this size`,
		},
		{
			Num:  "CommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Num:  "CommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deals in batch would start expiring`,
		},
		{
			Num:  "BatchPreCommitAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing precommit batching, instead
sending precommit messages to the chain individually`,
		},
		{
			Num:  "AggregateAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing commit aggregation, instead
submitting proofs to the chain individually`,
		},
		{
			Num:  "TerminateBatchMax",
			Type: "uint64",

			Comment: ``,
		},
		{
			Num:  "TerminateBatchMin",
			Type: "uint64",

			Comment: ``,
		},
		{
			Num:  "TerminateBatchWait",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Splitstore": []DocField{
		{
			Num:  "ColdStoreType",
			Type: "string",

			Comment: `ColdStoreType specifies the type of the coldstore.
It can be "universal" (default) or "discard" for discarding cold blocks.`,
		},
		{
			Num:  "HotStoreType",
			Type: "string",

			Comment: `HotStoreType specifies the type of the hotstore.
Only currently supported value is "badger".`,
		},
		{
			Num:  "MarkSetType",
			Type: "string",

			Comment: `MarkSetType specifies the type of the markset.
It can be "map" for in memory marking or "badger" (default) for on-disk marking.`,
		},
		{
			Num:  "HotStoreMessageRetention",
			Type: "uint64",

			Comment: `HotStoreMessageRetention specifies the retention policy for messages, in finalities beyond
the compaction boundary; default is 0.`,
		},
		{
			Num:  "HotStoreFullGCFrequency",
			Type: "uint64",

			Comment: `HotStoreFullGCFrequency specifies how often to perform a full (moving) GC on the hotstore.
A value of 0 disables, while a value 1 will do full GC in every compaction.
Default is 20 (about once a week).`,
		},
		{
			Num:  "EnableColdStoreAutoPrune",
			Type: "bool",

			Comment: `EnableColdStoreAutoPrune turns on compaction of the cold store i.e. pruning
where hotstore compaction occurs every finality epochs pruning happens every 3 finalities
Default is false`,
		},
		{
			Num:  "ColdStoreFullGCFrequency",
			Type: "uint64",

			Comment: `ColdStoreFullGCFrequency specifies how often to performa a full (moving) GC on the coldstore.
Only applies if auto prune is enabled.  A value of 0 disables while a value of 1 will do
full GC in every prune.
Default is 7 (about once every a week)`,
		},
		{
			Num:  "ColdStoreRetention",
			Type: "int64",

			Comment: `ColdStoreRetention specifies the retention policy for data reachable from the chain, in
finalities beyond the compaction boundary, default is 0, -1 retains everything`,
		},
	},
	"StorageMiner": []DocField{
		{
			Num:  "Subsystems",
			Type: "MinerSubsystemConfig",

			Comment: ``,
		},
		{
			Num:  "Dealmaking",
			Type: "DealmakingConfig",

			Comment: ``,
		},
		{
			Num:  "IndexProvider",
			Type: "IndexProviderConfig",

			Comment: ``,
		},
		{
			Num:  "Proving",
			Type: "ProvingConfig",

			Comment: ``,
		},
		{
			Num:  "Sealing",
			Type: "SealingConfig",

			Comment: ``,
		},
		{
			Num:  "Storage",
			Type: "SealerConfig",

			Comment: ``,
		},
		{
			Num:  "Fees",
			Type: "MinerFeeConfig",

			Comment: ``,
		},
		{
			Num:  "Addresses",
			Type: "MinerAddressConfig",

			Comment: ``,
		},
		{
			Num:  "DAGStore",
			Type: "DAGStoreConfig",

			Comment: ``,
		},
	},
	"Wallet": []DocField{
		{
			Num:  "RemoteBackend",
			Type: "string",

			Comment: ``,
		},
		{
			Num:  "EnableLedger",
			Type: "bool",

			Comment: ``,
		},
		{
			Num:  "DisableLocal",
			Type: "bool",

			Comment: ``,
		},
	},
}
