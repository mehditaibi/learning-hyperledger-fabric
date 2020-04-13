# Hyperledger Key Features 

## Key Features 

* Assets 
* Chaincode 
* Ledger Features 
* Privacy (Channels)
* Security and Membership Services 
* Consensus

# Assets 

Assets are represented in Hyperledeger Fabric as a collection of key-value pairs. All the changes in the state of assets are recorded as transactions on a Channel ledger where assets are defined. 

Assets can be tangible entities like hard drives, USB's, etc. to intangible objects like contracts, etc. 
Hyperledger Fabric provides the creation of assets using Chaincode and Hyperledger Composer tool.

Assets in Hyperledger can be represented in either binary or JSON format. 

# Chaincode 

* Chaincode is the other name for the smart contracts in Hyperledger. 
* Chaincode defines the business logic between multiple parties over the Hyperledger Blockchain network.
* Only chaincode can read or update data over the Hyperledger network. 
* Chaincode must be installed on every single peer of a channel or organization. 
* There can be multiple chaincodes with different policies.
* Chaincodes are of two types: 
    * Developed Chaincode - This is written in Go, NodeJS and Java programming language and deployed by users over the network to perfom business logic or specific functionalities. 
    * System Chaincode - This is used by the core of Hyperledger in managing the blockchain network to install, initiate and update the network entities. 

# Ledger Features 

The ledger is the sequenced, tamper-resistant record of all state transitions in the fabirc. State transistions are the resultant of the chaincode innovations submitted by participants.

A separate ledger is maintained per channel, and it stores the permanent, sequenced record in blocks, as 
well as a state database to maintain current fabric state. Each peer keeps a copy of the ledger for each channel of which they are a member. 

* Transactions are ordered into blocks over the ledger. 
* Transactions consist of verisons of key/value that are read from and written into chaincodes. 
* Peer validate transactions against endorsement policies and enfore the rules. 
* A channel's ledger contains a configuration block defining policies, access control lists, and other pertinent information.
* Query and update of the ledger is done by using key-based lookups, range queries and composite vital queries. 

# Privacy 

Hyperledger Fabric key benefit is the privacy which is provided though channels. Channels are similar to private network hosting their ledger where two or more members can conduct private and confidential transactions: 

* Hyperledger Fabric employs a ledger per channel
* Chaincodes are present to modify the state of ledger over the channels.
* A ledger exists in the scope of a channel. 
* Participants can connect to one or more channels in a Fabric network. 
* Data can further be obfuscated by encrypting the data before putting up a channel. 
* Channels provide fabrication of assets and participants over the Fabric network. 

# Security and Membership Services 

Hyperledger Fabric governs the participants with their identities. These identities are generated through a trusted Membership Service Provider (MSP).

* Public Key Infrastructure is used to generate cryptographic certificates which are tied to organizations, network components, and end users or client applications. 
* Data access control can be manipulated and governed on the whole network and channel levels, using the cryptographic certificates. 
* Pricacy and confidentiality can be maintained by combining channels and memebership services. 

# Consensus 

The consensus Hyperledger Fabric Network is a process where the nodes or computers in the network 
provide a guaranteed ordering of the transaction and valdation of block that needs to be commited to the ledger. 
Consensus must ensure the following in the network: 

* Confimr the correctness of all transactions in a proposed block according to endorsement and consensus policies. 
* Agree on order and correctness and hence on resuls of execution of the smart contract.
* Interface and depend on the smart-contract layer to verify the correctness of an ordered set of transactions in a block.

Consensus models in Hyperledger: 

* Permissionel-voting-based 
* Lottery based 