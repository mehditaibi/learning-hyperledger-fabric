# Steps Involved In Transaction Flow

## Transaction Proposal

* A client or an organisation uses Hyperledger SDK or an application interface to initiate the transaction. 
* Transaction proposal is prepared and is sent to endorsing peers. 
* The proposal is a request to invoke a chaincode function so that data can be read and/ or written to the peer ledger. 

# Endorsement Response 

* Enforsing peers check for the validity of the format of the transaction proposal, for the duplicity of the transaction proposal, the signature and authorization of the requesting client is checked using Membership Service Provider. 
* Endorsing peers take transaction proposal inputs as arguments and pass to chaincode function for simulation. 
* Chaincode gets executed against the current state database to produce response value, read set, and write set. 
* These sets along with the endorsing peer's signature is passed back as a "proposal response" to the SDK. 

# Verification of Endorsement 

* Application verifies endorsing peer signatures and comprares the endorsement response to determine if the proposal 
responses are the same. 
* If application only queried the ledger, then the query response is inspected. 
* Moreover, the application determines if the specified endorsement policy has been fulfilled before submitting the transaction to ordering service. 

# Broadcasting to Ordering Service 

* The application "broadcasts" the transaction proposal and response within a "transaction message" to the Ordering Service. 
* The transaction contains the read/write sets, the endorsing peers signatures and the Channel ID. 
* The Ordering Servie simply receives transactions from all channels in the network, orders them chronologically by channel, and creates blocks of transactions per channel. 
