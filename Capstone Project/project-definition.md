## Phase Definition 

We are going to perform the project in two phases: 

 1. Designing a network for organizations who are going to be part of the project. 
 2. Creating a chaincode for business logic.

## Designing the Network 

The first step in determining a Hyperledger Fabirc network strucutre for one's application is listing the participating organizations. Our network is going to have 4 organizations: 

 * Intel (Exporter)
 * Lenovo (Importer)
 * FedEx (Carrier)
 * GovRegulator (Regulator)

As we are crafting a sample applicationm that's why we are going to create two peers each for each organization. For production environment, you might have multiple peers. Apart from the peers, our network consists of one MSP for each of the four organizations, and an ordering service running in solo mode (a separate organization in itself).

## Steps in Designing our Network 

### Step 1: Generate network Cryptophic Material ( Certificate for the organizations)

### Step 2: Generate Channel Artifacts ( Transactions for Channel Creation, Genesis Block and Anchor Peers)

### Step 3: Generate the configuration for operating the network. (Specify peer docker images and capabilities)

### Step 4: Sarting the network