a = eth.accounts[0]
web3.eth.defaultAccount = a;

var abi = [
	{
		"inputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "dst",
				"type": "address"
			}
		],
		"name": "ping",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "pong",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]

var bytecode = "0x608060405234801561001057600080fd5b506040805180820190915260028082527f486900000000000000000000000000000000000000000000000000000000000060209092019182526100559160009161005b565b506100f6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009c57805160ff19168380011785556100c9565b828001600101855582156100c9579182015b828111156100c95782518255916020019190600101906100ae565b506100d59291506100d9565b5090565b6100f391905b808211156100d557600081556001016100df565b90565b6102bb806101056000396000f3fe608060405234801561001057600080fd5b5060043610610052577c0100000000000000000000000000000000000000000000000000000000600035046314df9e438114610057578063bc9748a1146100ff575b600080fd5b61008a6004803603602081101561006d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610107565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a6101f9565b60608173ffffffffffffffffffffffffffffffffffffffff1663bc9748a16040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160006040518083038186803b15801561016b57600080fd5b505afa15801561017f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156101a857600080fd5b8101908080516401000000008111156101c057600080fd5b820160208101848111156101d357600080fd5b81516401000000008111828201871017156101ed57600080fd5b50909695505050505050565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102855780601f1061025a57610100808354040283529160200191610285565b820191906000526020600020905b81548152906001019060200180831161026857829003601f168201915b505050505090509056fea165627a7a72305820e11f3eefecca94aa21a3bc7b2fa243ae61072217f9774dbcbb2fe5cde2a978560029";

var pingPong = web3.eth.contract(abi);
var private1 = pingPong.new({from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ['QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=']}, function(e, contract) {
	if (e) {
		console.log("err creating contract", e);
	} else {
		if (!contract.address) {
			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
		} else {
			console.log("Contract mined! Address: " + contract.address);
			console.log(contract);
		}
	}
});

var private2 = pingPong.new({from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ['QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=']}, function(e, contract) {
	if (e) {
		console.log("err creating contract", e);
	} else {
		if (!contract.address) {
			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
		} else {
			console.log("Contract mined! Address: " + contract.address);
			console.log(contract);
		}
	}
});

var public1 = pingPong.new({from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760 }, function(e, contract) {
	if (e) {
		console.log("err creating contract", e);
	} else {
		if (!contract.address) {
			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
		} else {
			console.log("Contract mined! Address: " + contract.address);
			console.log(contract);
		}
	}
});

var public2 = pingPong.new({from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760 }, function(e, contract) {
	if (e) {
		console.log("err creating contract", e);
	} else {
		if (!contract.address) {
			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
		} else {
			console.log("Contract mined! Address: " + contract.address);
			console.log(contract);
		}
	}
});
