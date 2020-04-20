# STEPS 

 1. Wrote the crypto-config.yaml file in the loc folder. 

 2. Copied cryptogen executable into the Capstone folder then ran: 
  ./cryptogen generate --config=./loc/crypto-config.yaml    
  The above generated the crypto-config folder. 

 3. Wrote the configtx.yaml file in the loc folder.

 4.  Running Configtxgen
  All the commands mentioned under this article are to be ran with the development unix environment where we have the Hyperledger Fabric tools.

  The following commands are going to create new folders and files under your working directory. Make sure you provide the proper permissions for the same.

 5. Before starting up the process to generate artifacts, we need to tell the configtxgen tool where to look for the configtx.yaml file that it needs to ingest. We will tell it look in our present working directory. We will add an environment variable, which will help us notify the configtxgen, run the following in your terminal:

export FABRIC_CFG_PATH=$PWD
After this we are going to create a folder where all the generated material will be stored. Let's add a new folder under your project directory:

sudo mkdir channel-artifacts

You also need to provide proper permissions to the channel-artifacts directory. As we are using a sample network, for now we are going to open all the permissions to the directory by using this command:

sudo chmod 777 channel-artifacts

6. Once we have setup the directory structure, we can go into generating the genesis block using the profile, which we created under the Configuration file, run the following command to generate the genesis block:

../bin/configtxgen -profile FourOrgsLoCGenesis -channelID loc-system-channel -outputBlock ./channel-artifacts/genesis.block

profile: This is the genesis profile which we created under the configuration.

channelID: This is the system channel ID, which will store all the configurations. This channel is not accessible by the applications. All the blocks under the channel are going to follow the configuration mentioned under the system channel. Make sure your application channel name is different from the system channel name.

outputBlock: This is the output file generated after the execution of Configtxgen.

7. If you want to generate genesis block for Kafka Ordering Service, then you can use the following command:

../bin/configtxgen -profile SampleDevModeKafka -channelID loc-sys-channel -outputBlock ./channel-artifacts/genesis.block
If you want to generate the genesis block for Raft Ordering Service, then you can use the following command:

../bin/configtxgen -profile SampleMultiNodeEtcdRaft -channelID loc-sys-channel -outputBlock ./channel-artifacts/genesis.block
Next, we need to create the channel transaction artifact. We are going to use again the profile defined under the Channel configuration file. Before generating the channel transaction, let's import another environment variable, such that we don't need to pass in the channel name again. To do that, let's run the following command:

export CHANNEL_NAME=loc-channel

8. Now, we can run the command to generate the channel transaction:

../bin/configtxgen -profile FourOrgsLoCChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

profile: This is the channel profile which we created under the configuration.

outputCreateChannelTx: This is the output channel transaction file generated after the execution of Configtxgen.

channelID: This is the application channel ID which will be used by the organizations to interact with each other.

Note that you don’t have to issue a special command for the channel if you are using a Raft or Kafka ordering service. The FourOrgsLoCChannel profile will use the ordering service configuration you specified when creating the genesis block for the network.

9. Next, we will define the anchor peer for Intel Organization on the channel that we are constructing. Anchor peer are defined under transaction, which will be put up on top of our channel.  Use the following command to setup Anchor Peer Transaction:

../bin/configtxgen -profile FourOrgsLoCChannel -outputAnchorPeersUpdate ./channel-artifacts/IntelMSPanchors.tx -channelID $CHANNEL_NAME -asOrg IntelMSP

profile: This is the channel profile which we created under the configuration.

outputAnchorPeersUpdate: This is the output transaction file generated after the execution through Configtxgen.

channelID: This is the application channel ID which will be used by the organizations to interact with each other.

asOrg: This states that we are running the command as the specified organization.

Now we can define the anchor peer transactions for all the other organizations.

Lenovo:

../bin/configtxgen -profile FourOrgsLoCChannel -outputAnchorPeersUpdate ./channel-artifacts/LenovoMSPanchors.tx -channelID $CHANNEL_NAME -asOrg LenovoMSP

FedEx:

../bin/configtxgen -profile FourOrgsLoCChannel -outputAnchorPeersUpdate ./channel-artifacts/FedExMSPanchors.tx -channelID $CHANNEL_NAME -asOrg FedExMSP

GovRegulator:

../bin/configtxgen -profile FourOrgsLoCChannel -outputAnchorPeersUpdate ./channel-artifacts/GovRegulatorMSPanchors.tx -channelID $CHANNEL_NAME -asOrg GovRegulatorMSP

You can read more about the different commands under configtxgen by following this link: https://hyperledger-fabric.readthedocs.io/en/release-1.4/commands/configtxgen.html


10. To create & start all the containers: 
docker-compose -f docker-compose-e2e.yaml up -d
docker-compose -f docker-compose-e2e.yaml down


