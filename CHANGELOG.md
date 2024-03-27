# Changelog

## UNRELEASED

### Bug Fixes

- [#1336](https://github.com/swag-eag/swa/pull/1336) Update ethermint to develop to fix feeHistory rpc api.

*February 28, 2024*

## v1.1.0-rc5

### Bug Fixes

- [#1329](https://github.com/swag-eag/swa/pull/1329) Update cosmos-sdk to `v0.47.10`.

*February 19, 2024*

## v1.1.0-rc4

### State Machine Breaking

- [#1318](https://github.com/swag-eag/swa/pull/1318) Add packet_sequence index in relayer event.
- [#1318](https://github.com/swag-eag/swa/pull/1318) Fix filter rule for eth_getLogs.
- [#1322](https://github.com/swag-eag/swa/pull/1322) Add `v1.1.0-testnet-1` upgrade plan for testnet.

### Improvements

- [#1324](https://github.com/swag-eag/swa/pull/1324) Update cosmos-sdk to `v0.47.9`.

*February 5, 2024*

## v1.1.0-rc3

### Bug Fixes

- [#1292](https://github.com/swag-eag/swa/pull/1292) memiavl cancel background snapshot rewriting when graceful shutdown.
- [#1294](https://github.com/swag-eag/swa/pull/1294) Update ethermint to fix and improve of debug_traceCall and eth_feeHistory.
- [#1302](https://github.com/swag-eag/swa/pull/1302) Fix concurrent map access in rootmulti store.
- [#1304](https://github.com/swag-eag/swa/pull/1304) Write versiondb with fsync, and relax the version requirement on startup.
- [#1308](https://github.com/swag-eag/swa/pull/1308) Update ethermint to fix duplicate cache events emitted from evm hooks and wrong priority tx.
- [#1311](https://github.com/swag-eag/swa/pull/1311) Add missing version in memiavl log.

### Improvements

- [#1291](https://github.com/swag-eag/swa/pull/1291) Update ibc-go to v7.3.2.
- [#1309](https://github.com/swag-eag/swa/pull/1309) Add missing destroy for file lock and close map on error.

*January 5, 2024*

## v1.1.0-rc2

- [#1258](https://github.com/swag-eag/swa/pull/1258) Support hard-fork style upgrades.
- [#1272](https://github.com/swag-eag/swa/pull/1272) Update ethermint to develop, cosmos-sdk to `v0.47.7`.
- [#1273](https://github.com/swag-eag/swa/pull/1273) Enable push0 opcode in integration test.
- [#1274](https://github.com/swag-eag/swa/pull/1274) Remove authz module.
- [#1287](https://github.com/swag-eag/swa/pull/1287) Support debug_traceCall.

### Bug Fixes

- [#1215](https://github.com/swag-eag/swa/pull/1215) Update ethermint to fix of concurrent write in fee history.
- [#1217](https://github.com/swag-eag/swa/pull/1217) Use the default chain-id behavour in sdk.
- [#1216](https://github.com/swag-eag/swa/pull/1216) Update ethermint to fix of avoid redundant parse chainID from gensis when start server.
- [#1230](https://github.com/swag-eag/swa/pull/1230) Fix mem store in versiondb multistore.
- [#1233](https://github.com/swag-eag/swa/pull/1233) Re-emit logs in callback contract.
- [#1256](https://github.com/swag-eag/swa/pull/1256) Improve permission checkings for some messages.

### State Machine Breaking

- [#1232](https://github.com/swag-eag/swa/pull/1232) Adjust require gas in relayer precompile to be closed with actual consumed.
- [#1209](https://github.com/swag-eag/swa/pull/1209) Support accurate estimate gas in evm tx from relayer.
- [#1247](https://github.com/swag-eag/swa/pull/1247) Update ethermint to develop, go-ethereum to `v1.11.2`.
- [#1235](https://github.com/swag-eag/swa/pull/1235) Add channel detail in ica packet callback.
- [#1251](https://github.com/swag-eag/swa/pull/1251) Adjust require gas for submitMsgs in ica precompile.
- [#1252](https://github.com/swag-eag/swa/pull/1252) Add plan `v1.1.0-testnet` to update default max_callback_gas param.

### Improvements

- [#1239](https://github.com/swag-eag/swa/pull/1239) Refactor websocket/subscription system to improve performance and stability.
- [#1241](https://github.com/swag-eag/swa/pull/1241) Improve parallelization of memiavl restoration.
- (deps) [#1253](https://github.com/swag-eag/swa/pull/1253) Upgrade Go-Ethereum version to [`v1.11.6`](https://github.com/ethereum/go-ethereum/releases/tag/v1.11.6).

*October 17, 2023*

## v1.1.0-rc1

### Bug Fixes

- [#1206](https://github.com/swag-eag/swa/pull/1206) Add missing keypair of SendEnabled to restore legacy param set before migration.
- [#1205](https://github.com/swag-eag/swa/pull/1205) Fix versiondb and memiavl upgrade issues, add integration test.

### Improvements

- [#1197](https://github.com/swag-eag/swa/pull/1197) tune rocksdb options to control memory consumption.

*October 9, 2023*

## v1.1.0-rc0

### State Machine Breaking

- [swa#695](https://github.com/swag-eag/swa/pull/695) Implement ADR-007, generic events format with indexed params.
- [swa#728](https://github.com/swag-eag/swa/pull/728) Upgrade gravity bridge latest bugfix, including multi attestation processing and double spend check.
- [swa#742](https://github.com/swag-eag/swa/pull/742) Add upgrade handler for v0.8.0-gravity-alpha2.
- [swa#750](https://github.com/swag-eag/swa/pull/750) Add upgrade handler for v0.8.0-gravity-alpha3.
- [swa#769](https://github.com/swag-eag/swa/pull/769) Prevent cancellation function to be called outside the scope of the contract that manage it.
- [swa#775](https://github.com/swag-eag/swa/pull/775) Support turnbridge transaction.
- [swa#781](https://github.com/swag-eag/swa/pull/781) Add prune command.
- [swa#830](https://github.com/swag-eag/swa/pull/830) Upgrade gravity bridge for latest bugfixes, patching two important DOS vulnerabilities
- [swa#834](https://github.com/swag-eag/swa/pull/834) Remove unsafe experimental flag.
- [swa#842](https://github.com/swag-eag/swa/pull/842) Add upgrade handler for v2.0.0-testnet3.
- [swa#795](https://github.com/swag-eag/swa/pull/795) Support permissions in swa.
- [swa#997](https://github.com/swag-eag/swa/pull/997) Fix logic to support proxy contract for swa originated swac20.
- [swa#1005](https://github.com/swag-eag/swa/pull/1005) Support specify channel id for send-to-ibc event in case of source token.
- [swa#1069](https://github.com/swag-eag/swa/pull/1069) Update ethermint to develop, go-ethereum to `v1.10.26` and ibc-go to `v6.2.0`.
- [swa#1147](https://github.com/swag-eag/swa/pull/1147) Integrate ica module.
- (deps) [#1121](https://github.com/swag-eag/swa/pull/1121) Bump Cosmos-SDK to v0.47.5 and ibc-go to v7.2.0.
- [swa#1014](https://github.com/swag-eag/swa/pull/1014) Support stateful precompiled contract for relayer.
- [swa#1165](https://github.com/swag-eag/swa/pull/1165) Icaauth module is not adjusted correctly in ibc-go v7.2.0.
- [swa#1163](https://github.com/swag-eag/swa/pull/1163) Support stateful precompiled contract for ica.
- [swa#837](https://github.com/swag-eag/swa/pull/837) Support stateful precompiled contract for bank.
- [swa#1184](https://github.com/swag-eag/swa/pull/1184) Update ibc-go to `v7.3.1`.
- [swa#1186](https://github.com/swag-eag/swa/pull/1186) Enlarge the max block gas limit in new version.
- [swa#1187](https://github.com/swag-eag/swa/pull/1187) Disable gravity module in app.
- [swa#1185](https://github.com/swag-eag/swa/pull/1185) Support ibc callback.
- [swa#1196](https://github.com/swag-eag/swa/pull/1196) Skip register stateful precompiled contract for bank.

### Bug Fixes

- [#833](https://github.com/swag-eag/swa/pull/833) Fix rollback command.
- [#945](https://github.com/swag-eag/swa/pull/945) Fix no handler exists for proposal type error when update-client due to wrong ibc route.
- [#1036](https://github.com/swag-eag/swa/pull/1036) Fix memiavl import memory leak.
- [#1038](https://github.com/swag-eag/swa/pull/1038) Update ibc-go to `v5.2.1`.
- [#1042](https://github.com/swag-eag/swa/pull/1042) Avoid channel get changed when concurrent subscribe happens ([ethermint commit](https://github.com/swag-eag/ethermint/commit/72bbe0a80dfd3c586868e2f0b4fbed72593c45bf)).
- [#1058](https://github.com/swag-eag/swa/pull/1058) Fix decode log for multi topics in websocket subscribe ([ethermint commit](https://github.com/swag-eag/ethermint/commit/2136ad029860c819942ad1836dd3f42585002233)).
- [#1062](https://github.com/swag-eag/swa/pull/1062) Update cometbft `v0.34.29` with several minor bug fixes and low-severity security-fixes.
- [#1075](https://github.com/swag-eag/swa/pull/1075) Add missing close in memiavl to avoid resource leaks.
- [#1073](https://github.com/swag-eag/swa/pull/1073) memiavl automatically truncate corrupted wal tail.
- [#1087](https://github.com/swag-eag/swa/pull/1087) memiavl fix LastCommitID when memiavl db not loaded.
- [#1088](https://github.com/swag-eag/swa/pull/1088) memiavl fix empty value in write-ahead-log replaying.
- [#1102](https://github.com/swag-eag/swa/pull/1102) avoid duplicate cache events emitted from ibc and gravity hook.
- [#1123](https://github.com/swag-eag/swa/pull/1123) Fix memiavl snapshot switching
- [#1125](https://github.com/swag-eag/swa/pull/1125) Fix genesis migrate for feeibc, evm, feemarket and gravity.
- [#1130](https://github.com/swag-eag/swa/pull/1130) Fix lock issues when state-sync with memiavl.
- [#1150](https://github.com/swag-eag/swa/pull/1150) Fix memiavl's unsafe retain of the root hashes.

### Features

- [#1042](https://github.com/swag-eag/swa/pull/1042) call Close method on app to cleanup resource on graceful shutdown ([ethermint commit](https://github.com/swag-eag/ethermint/commit/0ea7b86532a1144f229961f94b4524d5889e874d)).
- [#1083](https://github.com/swag-eag/swa/pull/1083) memiavl support both sdk 46 and 47 root hash rules.
- [#1091](https://github.com/swag-eag/swa/pull/1091) memiavl support rollback.
- [#1100](https://github.com/swag-eag/swa/pull/1100) memiavl support read-only mode, and grab exclusive lock for write mode.
- [#1103](https://github.com/swag-eag/swa/pull/1103) Add EventQueryTxFor cmd to subscribe and wait for transaction.
- [#1108](https://github.com/swag-eag/swa/pull/1108) versiondb support restore from local snapshot.
- [#1114](https://github.com/swag-eag/swa/pull/1114) memiavl support `CacheMultiStoreWithVersion`.
- [#1116](https://github.com/swag-eag/swa/pull/1116) versiondb commands support sdk47 app hash calculation.

### Improvements

- [#890](https://github.com/swag-eag/swa/pull/890) optimize memiavl snapshot format.
- [#904](https://github.com/swag-eag/swa/pull/904) Enable "dynamic-level-bytes" on new `application.db`.
- [#924](https://github.com/swag-eag/swa/pull/924) memiavl support `Export` API.
- [#950](https://github.com/swag-eag/swa/pull/950) Implement memiavl and integrate with state machine.
- [#985](https://github.com/swag-eag/swa/pull/985) Fix versiondb verify command on older versions
- [#998](https://github.com/swag-eag/swa/pull/998) Bump grocksdb to v1.7.16 and rocksdb to v7.10.2
- [#1028](https://github.com/swag-eag/swa/pull/1028) Add memiavl configs into app.toml
- [#1027](https://github.com/swag-eag/swa/pull/1027) Integrate local state-sync commands.
- [#1029](https://github.com/swag-eag/swa/pull/1029) Change config `async-commit` to `async-commit-buffer`, make the channel size configurable.
- [#1034](https://github.com/swag-eag/swa/pull/1034) Support memiavl snapshot strategy configuration.
- [#1035](https://github.com/swag-eag/swa/pull/1035) Support caching in memiavl directly, ignore inter-block cache silently.
- [#1050](https://github.com/swag-eag/swa/pull/1050) nativebyteorder mode will check endianness on startup, binaries are built with nativebyteorder by default.
- [#1064](https://github.com/swag-eag/swa/pull/1064) Simplify memiavl snapshot switching.
- [#1067](https://github.com/swag-eag/swa/pull/1067) memiavl: only export state-sync snapshots on an exist snapshot
- [#1082](https://github.com/swag-eag/swa/pull/1082) Make memiavl setup code reusable.
- [#1092](https://github.com/swag-eag/swa/pull/1092) memiavl disable sdk address cache if zero-copy enabled, and disable zero-copy by default.
- [#1099](https://github.com/swag-eag/swa/pull/1099) clean up memiavl tmp directories left behind.
- [#940](https://github.com/swag-eag/swa/pull/940) Update rocksdb dependency to 8.1.1.
- [#1149](https://github.com/swag-eag/swa/pull/1149) memiavl support `WorkingHash` api required by `FinalizeBlock`.
- [#1151](https://github.com/swag-eag/swa/pull/1151) memiavl `CacheMultiStoreWithVersion` supports `io.Closer`.
- [#1154](https://github.com/swag-eag/swa/pull/1154) Remove dependency on cosmos-sdk.
- [#1171](https://github.com/swag-eag/swa/pull/1171) Add memiavl background snapshot writing concurrency limit.
- [#1179](https://github.com/swag-eag/swa/pull/1179) Support blocking addresses in mempool.
- [#1182](https://github.com/swag-eag/swa/pull/1182) Bump librocksdb to 8.5.3.
- [#1183](https://github.com/swag-eag/swa/pull/1183) Avoid redundant logs added from relayer.

*April 13, 2023*

## v1.0.7

### Improvements

- [#936](https://github.com/swag-eag/swa/pull/936) Reuse recovered sender address to optimize performance ([ethermint commit](https://github.com/swag-eag/ethermint/commit/cb741e1d819683795aa32e286d31d8155f903cae)).
- [#949](https://github.com/swag-eag/swa/pull/949) Release static-linked binaries for linux platform.
- [#934](https://github.com/swag-eag/swa/pull/934) Add pebbledb backend.

### Bug Fixes

- [#953](https://github.com/swag-eag/swa/pull/953) Include third-party bug fixes:
  - update ethermint to include two bug fixes
    - <https://github.com/swag-eag/ethermint/pull/234>
    - <https://github.com/swag-eag/ethermint/pull/233>
  - update cosmos-sdk to include one bug fix
    - <https://github.com/cosmos/cosmos-sdk/pull/15667>
- [#945](https://github.com/swag-eag/swa/pull/945) Fix no handler exists for proposal type error when update-client due to wrong ibc route.

*Mar 16, 2023*

## v1.0.6

### Bug Fixes

- [#932](https://github.com/swag-eag/swa/pull/932) Backport multiple json-rpc bug fixes in ethermint ([commits](https://github.com/swag-eag/ethermint/compare/v0.20.8-swa...v0.20.9-swa)).

*Mar 6, 2023*

## v1.0.5

### Bug Fixes

- [#908](https://github.com/swag-eag/swa/pull/908) Forbids negative priority fee.

### Improvements

- [#904](https://github.com/swag-eag/swa/pull/904) Enable "dynamic-level-bytes" on new `application.db`.
- [#907](https://github.com/swag-eag/swa/pull/907) Apply a configurable limit in rpc apis.
- [#909](https://github.com/swag-eag/swa/pull/909) Update to cosmos-sdk v0.46.11.

*Feb 15, 2023*

## v1.0.4

### Bug Fixes

- [#814](https://github.com/swag-eag/swa/pull/814) Fix prometheus metrics.

### Improvements

- [#813](https://github.com/swag-eag/swa/pull/813) Tune up rocksdb options.
- [#791](https://github.com/swag-eag/swa/pull/791) Implement versiondb and migration commands.
- [#779](https://github.com/swag-eag/swa/pull/779) Add config iavl-lazy-loading to enable lazy loading of iavl store.

*Feb 08, 2023*

## v1.0.3

### Bug Fixes

- [#846](https://github.com/swag-eag/swa/pull/846) Disable authz message

*Jan 04, 2023*

## v1.0.2

### State Machine Breaking

- [#802](https://github.com/swag-eag/swa/pull/802) Update ibc-go to `v5.2.0`.

*December 14, 2022*

## v1.0.1

### Improvements

- [#781](https://github.com/swag-eag/swa/pull/781) Add prune command.
- [#790](https://github.com/swag-eag/swa/pull/790) Update cosmos-sdk to `v0.46.7`, it fix a migration issue which affects pending proposals's votes during upgrade,
  it also adds the config entries for file streamer.

*Nov 22, 2022*

## v1.0.0

### Improvements

- [#772](https://github.com/swag-eag/swa/pull/772) Update cosmos-sdk to `v0.46.6`, it's non-breaking for swa.

*Nov 17, 2022*

## v1.0.0-rc4

### Bug Fixes

- [#771](https://github.com/swag-eag/swa/pull/771) Fix london hardfork number in testnet3 parameters.

*Nov 13, 2022*

## v1.0.0-rc3

### State Machine Breaking

- [#765](https://github.com/swag-eag/swa/pull/765) Upgrade ibc-go to [v5.1.0](https://github.com/cosmos/ibc-go/releases/tag/v5.1.0) and related dependencies.

*Nov 10, 2022*

## v1.0.0-rc2

### Bug Fixes

- [#761](https://github.com/swag-eag/swa/pull/761) Fix non-deterministic evm execution result when there are concurrent grpc queries.
- [#762](https://github.com/swag-eag/swa/pull/762) Add `v1.0.0` upgrade plan for dry-run and mainnet upgrade, which clears the `extra_eips` parameter.
- [#763](https://github.com/swag-eag/swa/pull/763) Add error log for iavl set error.
- [#764](https://github.com/swag-eag/swa/pull/764) Make `eth_getProof` result compatible with ethereum.

*Nov 4, 2022*

## v1.0.0-rc1

### Bug Fixes

- [#760](https://github.com/swag-eag/swa/pull/760) Revert breaking changes on gas used in Ethermint.

*Nov 1, 2022*

## v1.0.0-rc0

### Bug Fixes

- [#748](https://github.com/swag-eag/swa/pull/748) Fix inconsistent state if upgrade migration commit is interrupted.
- [#752](https://github.com/swag-eag/swa/pull/752) Update iavl to `v0.19.4`.

*Oct 15, 2022*

## v0.9.0-beta4

### Bug Fixes

- [swa#719](https://github.com/swag-eag/swa/pull/719) Fix `eth_call` for legacy blocks (backport #713).

### Improvements

- [swa#720](https://github.com/swag-eag/swa/pull/720) Add option `iavl-disable-fastnode` to disable iavl fastnode indexing migration (backport #714).
- [swa#721](https://github.com/swag-eag/swa/pull/721) Integrate the file state streamer (backport #702).
- [swa#730](https://github.com/swag-eag/swa/pull/730) Update dependencies to recent versions (backport #729).

*Sep 20, 2022*

## v0.9.0-beta3

### Bug Fixes

- [swa#696](https://github.com/swag-eag/swa/pull/696) Fix json-rpc apis for legacy blocks.

*Aug 29, 2022*

## v0.9.0-beta2

*September 13, 2022*

## v0.9.0

### State Machine Breaking

- [swa#429](https://github.com/swag-eag/swa/pull/429) Update ethermint to main, ibc-go to v3.0.0, cosmos sdk to v0.45.4 and gravity to latest, remove v0.7.0 related upgradeHandler.
- [swa#532](https://github.com/swag-eag/swa/pull/532) Add SendtoChain and CancelSendToChain support from evm call.
- [swa#600](https://github.com/swag-eag/swa/pull/600) Implement bidirectional token mapping.
- [swa#611](https://github.com/swag-eag/swa/pull/611) Fix mistake on acknowledgement error in ibc middleware.
- [swa#627](https://github.com/swag-eag/swa/pull/627) Upgrade gravity bridge module with security enhancements
- [swa#647](https://github.com/swag-eag/swa/pull/647) Integrate ibc fee middleware.
- [swa#672](https://github.com/swag-eag/swa/pull/672) Revert interchain-accounts integration.

### Bug Fixes

- [swa#502](https://github.com/swag-eag/swa/pull/502) Fix failed tx are ignored in json-rpc apis.
- [swa#556](https://github.com/swag-eag/swa/pull/556) Bump gravity bridge module version to include bugfixes (including grpc endpoint)
- [swa#639](https://github.com/swag-eag/swa/pull/639) init and validate-genesis commands don't include experimental modules by default.

### Improvements

- [swa#418](https://github.com/swag-eag/swa/pull/418) Support logs in evm-hooks and return id for SendToEthereum events
- [swa#489](https://github.com/swag-eag/swa/pull/489) Enable jemalloc memory allocator, and update rocksdb src to `v6.29.5`.
- [swa#511](https://github.com/swag-eag/swa/pull/511) Replace ibc-hook with ibc middleware, use ibc-go upstream version.
- [swa#550](https://github.com/swag-eag/swa/pull/550) Support basic json-rpc apis on pruned nodes.
- [swa#549](https://github.com/swag-eag/swa/pull/549) Use custom tx indexer feature of ethermint.
- [swa#673](https://github.com/swag-eag/swa/pull/673) Upgrade cosmos-sdk to 0.46.1 and ibc-go to v5.0.0-rc0.

*Aug 5, 2022*

## v0.8.0

### State Machine Breaking

- [swa#618](https://github.com/swag-eag/swa/pull/618) selfdestruct don't delete bytecode of smart contract.

*Aug 5, 2022*

## v0.7.1

### Bug Fixes

- [swa#454](https://github.com/swag-eag/swa/pull/454) Add back the latest testnet upgrade handler.
- [swa#503](https://github.com/swag-eag/swa/pull/503) Fix failed tx are ignored in json-rpc apis (backport #502).
- [swa#526](https://github.com/swag-eag/swa/pull/526) Fix tendermint duplicated tx issue.
- [swa#584](https://github.com/swag-eag/swa/pull/584) Validate eth tx hash in ante handler and fix tx hashes returned in some JSON-RPC apis.
- [swa#587](https://github.com/swag-eag/swa/pull/587) Unlucky tx patch cmd recompute eth tx hash.
- [swa#595](https://github.com/swag-eag/swa/pull/595) Workaround the tx hash issue in event parsing.

### Improvements

- [swa#489](https://github.com/swag-eag/swa/pull/489) Enable jemalloc memory allocator, and update rocksdb src to `v6.29.5`.
- [swa#513](https://github.com/swag-eag/swa/pull/513) Add `fix-unlucky-tx` command to patch txs post v0.7.0 upgrade.
- [swa#522](https://github.com/swag-eag/swa/pull/522) Add `reindex-duplicated-tx` command to handle the tendermint tx duplicated issue.
- [swa#585](https://github.com/swag-eag/swa/pull/585) Reject replay unprotected tx, mainly the old transactions on ethereum.

*May 3, 2022*

## v0.7.0

### State Machine Breaking

- [swa#241](https://github.com/swag-eag/swa/pull/241) Update ethermint to main and merged statedb refactoring in custom fork.
- [swa#289](https://github.com/swag-eag/swa/pull/289) Update ethermint to `v0.10.0-swa` which uses ibc-go `v2.0.2` instead of `v3.0.0-alpha2` and include the fixes below:
  - [ethermint#901](https://github.com/tharsis/ethermint/pull/901) support batch evm tx
  - [ethermint#849](https://github.com/tharsis/ethermint/pull/849) Change EVM hook interface.
  - [ethermint#809](https://github.com/tharsis/ethermint/pull/809) fix nonce increment issue when contract deployment tx get reverted.
  - [ethermint#855](https://github.com/tharsis/ethermint/pull/855) unify base fee related logic in the code.
  - [ethermint#817](https://github.com/tharsis/ethermint/pull/817) Fix eip-1559 logic related to effectiveGasPrice.
  - [ethermint#822](https://github.com/tharsis/ethermint/pull/822) Update base fee in begin blocker rather than end blocker.
  - [cosmos-sdk#10833](https://github.com/cosmos/cosmos-sdk/pull/10833) fix reported tx gas used when block gas limit exceeded.
  - [cosmos-sdk#10814](https://github.com/cosmos/cosmos-sdk/pull/10814) revert tx when block gas limit exceeded.
  - [cosmos-sdk#10725](https://github.com/cosmos/cosmos-sdk/pull/10725) populate `ctx.ConsensusParams` for begin/end blockers (fix baseFee calculation in ethermint).
- [swa#315](https://github.com/swag-eag/swa/pull/315) Update cosmos-sdk to `v0.45.0`

### Improvements

- [swa#210](https://github.com/swag-eag/swa/pull/210) re-enabling gravity bridge conditionally
- [swa#322](https://github.com/swag-eag/swa/pull/322) Merge min-gas-price change in ethermint: don't check min-gas-price for EVM tx when feemarket enabled.
- [swa#345](https://github.com/swag-eag/swa/pull/345) disable the url query parameter in swagger-ui.
- [swa#328](https://github.com/swag-eag/swa/pull/328) display detail panic information in query result when `--trace` enabled.
- [swa#441](https://github.com/swag-eag/swa/pull/441) Update cosmos-sdk to `v0.45.4`

### Bug Fixes

- [swa#287](https://github.com/swag-eag/swa/pull/287) call upgrade handler before sealing app
- [swa#323](https://github.com/swag-eag/swa/pull/323) Upgrade gravity bridge to v0.3.9 which contain a bugfix on `batchTxExecuted.`
- [swa#324](https://github.com/swag-eag/swa/pull/324) Update to cosmos-sdk `v0.45.1`, which fixes an OOM issue.
- [swa#329](https://github.com/swag-eag/swa/pull/329) Fix panic of eth_call on blocks prior to upgrade.
- [swa#340](https://github.com/swag-eag/swa/pull/340) Update dependencies to include several bug fixes: a) fix subscription deadlock issue in ethermint, b) fix data races `traceContext`.
- [swa#370](https://github.com/swag-eag/swa/pull/370) Update ethermint to fix a websocket bug, add websockets integration tests.
- [swa#378](https://github.com/swag-eag/swa/pull/378) Backport recent ethermint bug fixes: a) fix tx inclusion issue by report correct gasWanted of eth tx, b) Add buffer to eth_gasPrice response to fix client UX, c) Quick fix for eth_feeHistory when reward is nil, d) add returnValue message on tracing.
- [swa#446](https://github.com/swag-eag/swa/pull/446) Fix failure of query legacy block after upgrade.

*December 10, 2021*

## v0.6.5

### Bug Fixes

- [swa#255](https://github.com/swag-eag/swa/pull/255) fix empty topics in non-breaking way
- [swa#270](https://github.com/swag-eag/swa/pull/270) reject MsgEthereumTx wrapping tx without the extension option.

*November 30, 2021*

## v0.6.4

### Bug Fixes

- [swag-eag/ethermint#19](https://github.com/swag-eag/ethermint/pull/19) revert tharsis#786 because it contains consensus breaking changes

*November 29, 2021*

## v0.6.3

### Bug Fixes

- [tharsis#781](https://github.com/tharsis/ethermint/pull/781) fix empty transactions in getBlock
- [swag-eag/ethermint#15](https://github.com/swag-eag/ethermint/pull/15) web3 rpc api returns wrong block gas limit
- [swag-eag/ethermint#16](https://github.com/swag-eag/ethermint/pull/16) fix unwrap context panic in BlockMaxGasFromConsensusParams

### Improvements

- [tharsis#786](https://github.com/tharsis/ethermint/pull/786) Improve error message of `SendTransaction`/`SendRawTransaction` JSON-RPC APIs.
- [swa#222](https://github.com/swag-eag/swa/pull/222) change solc 0.6.11 to 0.6.8 (from dapp cachix) and update hermes to 0.8.

*November 19, 2021*

## v0.6.2

### Bug Fixes

- [tharsis#720](https://github.com/tharsis/ethermint/pull/720) traceTransaction fails for succesful tx
- [tharsis#743](https://github.com/tharsis/ethermint/pull/743) missing debug_tranceBlockByHash RPC method and fix debug_traceBlock*
- [tharsis#746](https://github.com/tharsis/ethermint/pull/746) set debug based on tracer
- [tharsis#741](https://github.com/tharsis/ethermint/pull/741) filter non eth txs in block rpc response
- [swag-eag/ethermint#12](https://github.com/swag-eag/ethermint/pull/12) reject tx with too large gas limit

*October 26, 2021*

## v0.6.1

### State Machine Breaking

- [swa#190](https://github.com/swag-eag/swa/pull/190) upgrade ethermint to v0.7.2 with (#661) and (#689)

### Bug Fixes

- [swa#187](https://github.com/swag-eag/swa/pull/187) multiple denoms can be mapped to same contract
- [swa#157](https://github.com/swag-eag/swa/pull/185) swa params name has an unnecessary Key prefix
- [swa#179](https://github.com/swag-eag/swa/pull/179) fix denom (symbol) in SWAC20Module
- [swa#178](https://github.com/swag-eag/swa/pull/178) version CLI command doesn't output any text

*October 13, 2021*

## v0.6.0

This version removes gravity-bridge from swa, also includes multiple bug fixes in third-party dependencies.

### Consensus breaking changes

- [swa#171](https://github.com/swag-eag/swa/pull/171) remove gravity-bridge for mainnet launch

### Bug Fixes

- [swa#144](https://github.com/swag-eag/swa/pull/144) fix events in autodeploy swac20 module contract
- [gravity-bridge#17](https://github.com/swag-eag/gravity-bridge/pull/17) processEthereumEvent does not persist hooks emitted event
- [gravity-bridge#20](https://github.com/swag-eag/gravity-bridge/pull/20) fix undeterministic in consensus
- [swa#167](https://github.com/swag-eag/swa/pull/167) upgrade cosmos-sdk to 0.44.2

### Improvements

- [swa#162](https://github.com/swag-eag/swa/pull/162) bump ibc-go to v1.2.1 with hooks support
- [swa#169](https://github.com/swag-eag/swa/pull/169) bump ethermint to v0.7.1 and go-ethereum to v10.1.3-patched which include (CVE-2021-39137) hotfix

*October 4, 2021*

## v0.5.5

This version fixes various bugs regarding ibc fund transfer and EVM-related in ethermint.
We also enable swagger doc ui and add the token mapping state in genesis.

### Bug Fixes

- [swa#109](https://github.com/swag-eag/swa/issues/109) ibc transfer timeout too short
- [tharsis#590](https://github.com/tharsis/ethermint/pull/590) fix export contract state in genesis and reimport
- [swa#123](https://github.com/swag-eag/swa/issues/123) fix ibc refund logic
- [tharsis#617](https://github.com/tharsis/ethermint/pull/617) iterator on deeply nested cache contexts is extremely slow
- [tharsis#615](https://github.com/tharsis/ethermint/pull/615) tx log attribtue value not parsable by some client

### Features

- [swa#110](https://github.com/swag-eag/swa/pull/110) embed swagger doc ui
- [swa#113](https://github.com/swag-eag/swa/pull/113) export token mapping state to genesis
- [swa#128](https://github.com/swag-eag/swa/pull/128) add native message to update token mapping

*September 22, 2021*

## v0.5.4

This version is the same as v0.5.3 with a patched version of ethermint which include a bug fix on the transaction receipts events and on concurrent query.

### Bug Fixes

- [swa#93](https://github.com/swag-eag/swa/pull/93) tx receipts don't contain events
- [swa#98](https://github.com/swag-eag/swa/pull/98) node crash under concurrent query

*September 21, 2021*

## v0.5.3

This version contains several new features, it enables gravity bridge in Swa and automatic token conversion for bridging tokens to swac20 tokens. It also fixes the decimal conversion issues in the CRO tokens from Crypto.org Chain.
In addition to that, it also upgrade ethermint to its latest version (v0.5.0.x) which bring several breaking changes (see [changelog](https://github.com/tharsis/ethermint/blob/1a01c6a992c0fb94d70bb1c7127715874cefd057/CHANGELOG.md)).

### Consensus breaking changes

- [swa#87](https://github.com/swag-eag/swa/pull/87) upgrade ethermint to v0.4.2-0.20210920104419-1a01c6a992c0

### Features

- [swa#11](https://github.com/swag-eag/swa/pull/11) embed gravity bridge module
- [swa#35](https://github.com/swag-eag/swa/pull/35) add support for ibc hook
- [swa#55](https://github.com/swag-eag/swa/pull/55) add support for ibc token conversion to swac20
- [swa#45](https://github.com/swag-eag/swa/pull/45) allow evm contract to call bank send and gravity send
- [swa#65](https://github.com/swag-eag/swa/pull/65) support SendToIbc in evm_log_handlers
- [swa#59](https://github.com/swag-eag/swa/pull/59) gravity bridged tokens are converted to swac20
  automatically
- [swa#68](https://github.com/swag-eag/swa/issues/68) support SendCroToIbc in evm_log_handlers
- [swa#86](https://github.com/swag-eag/swa/issues/86) change account prefix

*August 19, 2021*

## v0.5.2

### Consensus breaking changes

- (ethermint) [tharsis#447](https://github.com/tharsis/ethermint/pull/447) update `chain-id` format.

### Improvements

- (ethermint) [tharsis#434](https://github.com/tharsis/ethermint/pull/434) configurable vm tracer

### Bug Fixes

- (ethermint) [tharsis#446](https://github.com/tharsis/ethermint/pull/446) fix chain state export issue

*August 16, 2021*

## v0.5.1

This version is a new scaffolding of swa project where ethermint is included as a library.

### Consensus breaking changes

- (ethermint) [tharsis#399](https://github.com/tharsis/ethermint/pull/399) Exception in sub-message call reverts the call if it's not propagated.
- (ethermint) [tharsis#334](https://github.com/tharsis/ethermint/pull/334) Log index changed to the index in block rather than tx.
- (ethermint) [tharsis#342](https://github.com/tharsis/ethermint/issues/342) Don't clear balance when resetting the account.
- (ethermint) [tharsis#383](https://github.com/tharsis/ethermint/pull/383) `GetCommittedState` use the original context.

### Features

### Improvements

- (ethermint) [tharsis#425](https://github.com/tharsis/ethermint/pull/425) Support build on linux arm64
- (ethermint) [tharsis#423](https://github.com/tharsis/ethermint/pull/423) Bump to cosmos-sdk 0.43.0

### Bug Fixes

- (ethermint) [tharsis#428](https://github.com/tharsis/ethermint/pull/428) [tharsis#375](https://github.com/tharsis/ethermint/pull/375) Multiple web3 rpc api fixes.
