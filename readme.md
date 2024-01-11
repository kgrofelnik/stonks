# stonks
**stonks** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started


```
export OPENAI_API_KEY="your_key_here"
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Use

```
txhash=$(stonksd tx stonks ask "What is a grapefruit?" --from alice --chain-id stonks | grep "txhash:" | awk '{print $2}')
stonksd query tx $txhash --chain-id stonks
```