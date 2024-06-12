# LocalOsmosis

LocalOsmosis is a complete Osmosis testnet containerized with Docker and orchestrated with a simple docker-compose file. LocalOsmosis comes preconfigured with opinionated, sensible defaults for a standard testing environment.

LocalOsmosis comes in two flavors:

1. No initial state: brand new testnet with no initial state. 
2. With mainnet state: creates a testnet from a mainnet state export

Both ways, the chain-id for LocalOsmosis is set to 'localosmosis'.

## Prerequisites

Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

## 1. LocalOsmosis - No Initial State

The following commands must be executed from the root folder of the Osmosis repository.

1. Make any change to the osmosis code that you want to test

2. Initialize LocalOsmosis:

```bash
make localnet-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.osmosisd-local` folder

3. Start LocalOsmosis:

```bash
make localnet-start
```

> Note
>
> You can also start LocalOsmosis in detach mode with:
>
> `make localnet-startd`

4. (optional) Add your validator wallet and 9 other preloaded wallets automatically:

```bash
make localnet-keys
```

- These keys are added to your `--keyring-backend test`
- If the keys are already on your keyring, you will get an `"Error: aborted"`
- Ensure you use the name of the account as listed in the table below, as well as ensure you append the `--keyring-backend test` to your txs
- Example: `osmosisd tx bank send lo-test2 symphony1jpr5824frn5472qm73ckfe2c3rh6vrn4lvlgj7 --keyring-backend test --chain-id localosmosis`

5. You can stop chain, keeping the state with

```bash
make localnet-stop
```

6. When you are done you can clean up the environment with:

```bash
make localnet-clean
```

## 2. LocalOsmosis - With Mainnet State

Running LocalOsmosis with mainnet state is resource intensive and can take a bit of time.
It is recommended to only use this method if you are testing a new feature that must be thoroughly tested before pushing to production.

A few things to note before getting started. The below method will only work if you are using the same version as mainnet. In other words,
if mainnet is on v8.0.0 and you try to do this on a v9.0.0 tag or on main, you will run into an error when initializing the genesis.
(yes, it is possible to create a state exported testnet on a upcoming release, but that is out of the scope of this tutorial)

Additionally, this process requires 64GB of RAM. If you do not have 64GB of RAM, you will get an OOM error.

### Create a mainnet state export

1. Set up a node on mainnet (easiest to use the [get.osmosis.zone](https://get.osmosis.zone) tool). This will be the node you use to run the state exported testnet, so ensure it has at least 64GB of RAM.

```sh
curl -sL https://get.osmosis.zone/install > i.py && python3 i.py
```

2. Once the installer is done, ensure your node is hitting blocks.

```sh
source ~/.profile
journalctl -u osmosisd.service -f
```

3. Stop your Osmosis daemon

```sh
systemctl stop osmosisd.service
```

4. Take a state export snapshot with the following command:

```sh
cd $HOME
osmosisd export 2> state_export.json
```

After a while (~15 minutes), this will create a file called `state_export.json` which is a snapshot of the current mainnet state.

### Use the state export in LocalOsmosis

1. Copy the `state_export.json` to the `LocalOsmosis/state_export` folder within the osmosis repo

```sh
cp $HOME/state_export.json $HOME/osmosis/tests/LocalOsmosis/state_export/
```

6. Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

7. Build the `local:osmosis` docker image:

```bash
make localnet-state-export-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.osmosisd` folder

3. Start LocalOsmosis:

```bash
make localnet-state-export-start
```

> Note
>
> You can also start LocalOsmosis in detach mode with:
>
> `make localnet-state-export-startd`

When running this command for the first time, `local:osmosis` will:

- Modify the provided `state_export.json` to create a new state suitable for a testnet
- Start the chain

You will then go through the genesis initialization process. This will take ~15 minutes.
You will then hit the first block (not block 1, but the block number after your snapshot was taken), and then you will just see a bunch of p2p error logs with some KV store logs.
**This will happen for about 1 hour**, and then you will finally hit blocks at a normal pace.

9. On your host machine, add this specific wallet which holds a large amount of osmo funds

```sh
MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"
echo $MNEMONIC | osmosisd keys add wallet --recover --keyring-backend test
```

You now are running a validator with a majority of the voting power with the same state as mainnet state (at the time you took the snapshot)

10. On your host machine, you can now query the state-exported testnet:

```sh
osmosisd status
```

11. Here is an example command to ensure complete understanding:

```sh
osmosisd tx bank send wallet symphony1jpr5824frn5472qm73ckfe2c3rh6vrn4lvlgj7 10000000uosmo --chain-id testing1 --keyring-backend test
```

12. You can stop chain, keeping the state with

```bash
make localnet-state-export-stop
```

13. When you are done you can clean up the environment with:

```bash
make localnet-state-export-clean
```

Note: At some point, all the validators (except yours) will get jailed at the same block due to them being offline.

When this happens, it may take a little bit of time to process. Once all validators are jailed, you will continue to hit blocks as you did before.
If you are only running the validator for a short time (< 24 hours) you will not experience this.

## LocalOsmosis Accounts

LocalOsmosis is pre-configured with one validator and 9 accounts with ION and OSMO balances.

| Account   | Address                                            | Mnemonic                                                                                                                                                                   |
|-----------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| lo-val    | `symphony1n7eq7r0ktw5q3w76854f3hqtz2mu9h2p54jsyr`  | `twist record public tattoo eternal blame twin saddle control monster quarter laptop lizard insect excite lawn lake caution tumble promote grow expand flight unknown`                    |
| lo-test1  | `symphony1c605nvcw94rvvehrcdfj85qe09ulseyt0efhk7`  | `summer solar wall song wall defy cement stable punch dilemma proud absurd pizza metal surge fury brush clown sponsor april bounce thought stumble buzz`                       |
| lo-test2  | `symphony1jpr5824frn5472qm73ckfe2c3rh6vrn4lvlgj7`  | `opinion coin approve cause element gallery film party laugh fire stumble switch analyst burden labor welcome sketch inflict three include chapter document jungle fossil`              |
| lo-test3  | `symphony1amr6zrvs0hymf62qd5mwvshx94ul8cgfu9jtxn`  | `early float super rebel long paper inside favorite brown casual original youth spoil sure knee body abstract merry dragon memory either rack tide road`        |
| lo-test4  | `symphony1egts9ayaqr6t54ahs62awmz5smuf764uu5f5xv`  | `you man shell elder fix among swarm shrug same absent faith axis reunion adult album gauge speed tumble brick false story board marble toast` |
| lo-test5  | `symphony1450weujlqvtd0d5z59v388jmzwyk3e6qhlj5r5`  | `mass embark injury obvious usage what foil fat diary assist wise agent rigid renew have unlock hole record urge news uncover acoustic measure valid`        |
| lo-test6  | `symphony12mdnm5yv5dfz37qsu0eu60x8qwxxl0x7sqqzn0`  | `tongue key pattern victory pumpkin divorce during dizzy pond wonder pelican visa arrive pave tonight cheese van junk bounce ability length hidden battle garlic`                  |
| lo-test7  | `symphony1ar8mfrrtkwlm62wgu88d0cfleng5gl8y062gsn`  | `fork tourist chalk photo wrist entire shell barrel decide rigid trip fancy ostrich frozen table obscure velvet car manual similar kingdom pill visual sunset`                       |
| lo-test8  | `symphony1kvgujs5yg9h6l6e265smwx99fmnnmc4af5v0ah`  | `bar assist play fade chuckle machine cross history flock water panther this repeat category snack viable awesome pause announce leisure pony organ orient visa`                 |
| lo-test9  | `symphony1ww5e3y7ptw8h3lc0cumxe5lmcu3m53dn7qyn4k`  | `seed knife aunt flip mansion bird trophy cluster truth purpose weekend exact jewel beach visit stairs rich capable employ float omit harvest shell garment`       |
| lo-test10 | `symphony1tsehv6f0v7ce4gy7574thxnp6v8jx7jm4evkpe`  | `page charge below sponsor primary page comic arrest ancient express polar frozen interest ship slim label giant remain evidence monitor large maximum cage curious`     |

## Tests

### Software-upgrade test

To test a software upgrade, you can use the `submit_upgrade_proposal.sh` script located in the `scripts/` folder. This script automatically creates a proposal to upgrade the software to the specified version and votes "yes" on the proposal. Once the proposal passes and the upgrade height is reached, you can update your localosmosis instance to use the new version.

#### Usage 

To use the script:

1. make sure you have a running LocalOsmosis instance

2. run the following command:

```bash
./scripts/submit_upgrade_proposal.sh <upgrade version>
```

Replace `<upgrade version>` with the version of the software you want to upgrade to, for example. If no version is specified, the script will default to `v15` version.

The script does the following:

- Creates an upgrade proposal with the specified version and description.
- Votes "yes" on the proposal.

#### Upgrade

Once the upgrade height is reached, you need to update your `localosmosis` instance to use the new software. 

There are two ways to do this:

1. Change the image in the `docker-compose.yml` file to use the new version, and then restart LocalOsmosis using `make localnet-start`. For example:

```yaml
services:
  osmosisd:
    image: <NEW_IMAGE_I_WANT_TO_USE>
    # All this needs to be commented to don't build the image with local changes
    # 
    # build:
    #     context: ../../
    #     dockerfile: Dockerfile
    #     args:
    #     RUNNER_IMAGE: alpine:3.17
    #     GO_VERSION: 1.20
```

2. Checkout the Osmosis repository to a different `ref` that includes the new version, and then rebuild and restart LocalOsmosis using `make localnet-start`. Make sure to don't delete your `~/.osmosisd-local` folder.

## FAQ

Q: How do I enable pprof server in localosmosis?

A: everything but the Dockerfile is already configured. Since we use a production Dockerfile in localosmosis, we don't want to expose the pprof server there by default. As a result, if you would like to use pprof, make sure to add `EXPOSE 6060` to the Dockerfile and rebuild the localosmosis image.
