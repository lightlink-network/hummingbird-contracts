#!/bin/bash -i

# Example usage:
# ./testnet_up.sh https://arbitrum-sepolia.infura.io/v3/your-api-key --persist 

# Vars
forkUrl=$1
persist=false
persistLocation=".testnet.state.log"

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    --persist)
      persist=true
      shift
      ;;
    --help)
      echo "Usage: ./testnet_up.sh <fork-url> [--persist]"
      echo "  --persist: Persist the state of the testnet"
      exit 0
      ;;
    *)
      shift
      ;;
  esac
done


# Check dependencies
hash anvil 
if [ $? -eq 1 ]; then
  echo "Error: Anvil is not installed. Please install Anvil via foundry 'https://book.getfoundry.sh/getting-started/installation'"
  exit 1
fi

hash npm
if [ $? -eq 1 ]; then
  echo "Error: npm is not installed. Please install npm via 'https://www.npmjs.com/get-npm'"
  exit 1
fi

echo "Starting testnet setup..."
# Add any necessary setup commands below

echo "Generating accounts..."
npx hardhat run scripts/tools/accounts.ts > testnet.accounts.log
if [ $? -eq 1 ]; then
  echo "Error: Failed to generate accounts."
  exit 1
fi

echo "Starting Anvil testnet..."
echo " - Chain ID: 1337"
echo " - Port: 8545"
if [ $persist = true ]; then
  echo " - Persisting state"
  anvil --chain-id 1337 --fork-url $forkUrl --state $persistLocation &
else
  echo " - Not persisting state"
  anvil --chain-id 1337 --fork-url $forkUrl &
fi
ANVIL_PID=$!

function cleanup {
  echo "Cleaning up..."
  if kill -0 $ANVIL_PID; then
    echo "Stopping Anvil..."
    kill $ANVIL_PID
  fi
}

trap cleanup EXIT

echo "Waiting for Anvil to start..."
sleep 5
# Check if Anvil is running
if ! kill -0 $ANVIL_PID; then
  echo "Error: Anvil failed to start."
  exit 1
fi

echo "Deploying contracts..."
# only run if persist false or persistLocation doesnt exists 
if [[ $persist != true ]] || [[ ! -f $persistLocation ]]; then
  npx hardhat run scripts/deploy/deploy.ts --network localhost > testnet.deployments.log
  if [ $? -eq 1 ]; then
    echo "Error: Failed to deploy contracts."
    kill $ANVIL_PID
    exit 1
  fi
else 
  echo "Contracts already deployed, skipping..."
fi

if [ $? -eq 0 ]; then
  echo ""
  echo "🙌 Testnet setup complete!"
  echo "- accounts: ./testnet.accounts.log"
  echo "- deployments: ./testnet.deployments.log"
  if [ $persist = true ]; then
    echo "- state: ./$persistLocation"
  fi
  echo ""
  echo "🟢 Network is running. Press Ctrl+C to stop."
  echo ""
  wait $ANVIL_PID
else
  echo "Testnet setup failed."
  kill $ANVIL_PID
fi



