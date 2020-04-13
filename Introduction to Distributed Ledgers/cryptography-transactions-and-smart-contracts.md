# Cryptography, Transactions and Smart Contracts 

## Cryptography 

Blockchain utilizes cryptography in the core: 

* Asymmetric Cryptography algorithm like ECC is used for the generation of keys. (public & private keys)
* Hasing algorithm SHA256 is used to generate digital signatures. 
* RIPEMD160 is utilized for hashing addresses. 
* Networking protocols like TLS or Gossip Protocol are utilised for creating the peer to peer network.

## Transactions 

* Transactions are records of data in chronological order. 
* transactiosn are stored in a Merkle tree inside the Block.
* The transactions, when submitted, are picked up by the blockchain network and is inserted into a 'pool of unconfirmed transactions'. The transaction pool is a collection of all the transactions on that network that have been comfirmed yet. 
* Miners on the network select transactions from this pool and add them to their 'block'.
* transactions also contain metadata information which can be utilized to store data over the Blockchain. 

## Smart Contracts 

* Smart Contracts are the digital contracts signed between two parties and store over the immutavle ledger.
* Smart contracts help you exchange money, property, shares, or anything of value in a transparentm conflict-free way while avoiding the services of a middleman. 
* Contracts can be encoded on a any blockchain, but Ethereum is mostly used since it gives an unlimited processing capability.
* Hyperledger is also providing chain codes which are very similar to Smart Contracts.
* Example: Renting an apartment.