# Hyperledger Fabric Transactions 

All the udpates over the ledger performed through transactions. A transaction is invoked when we are calling a function under the chaincode. Transactions are the major entities that deliver the actual data for clients and applications. 
With Hyperledger Fabric, we have 2 states of transactions: 

* Invoke Transactions: These transactions are use to invole a function under chaincode, which change the state of ledger. 
* Deploy Transactions: These transactions are used to deploy chaincode over the peer, ledger or network. 

# Transaction Flow Steps

* Transaction proposal to the endorsing peers 
* Endorsement response by the endorsing peers 
* Verification of Endorsement response 
* Invocation request to the ordering services 
* Invocation response to the peers by validating and commiting the transactions 
* Ledger gets updated 

# Transaction Flow Assumptions 

* Channel is properly set up and running. 
* The application user has registered and enrolled with the organization's certificate authority (CA) and received back 
the necessary cryptographic material, which is used to authenticate to the network.
* The chaincode is installed on the peers and instantiated on the channel.
* The chaincode contains logic defining a set of transaction instructions and the agreed upon price for an asset. 
* An endorsement policy has also been set for this chaincode, stating that all the peers must endorese any transaction.