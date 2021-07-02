use crate::query::get_last_event_nonce_for_validator;
use deep_space::error::CosmosGrpcError;
use deep_space::Address as CosmosAddress;
use deep_space::Contact;
use gravity_proto::gravity::query_client::QueryClient as GravityQueryClient;
use gravity_proto::cosmos_sdk_proto::cosmos::base::abci::v1beta1::TxResponse;
use gravity_utils::get_with_retry::RETRY_TIME;
use std::time::{Duration, Instant};
use tokio::time::sleep;
use tonic::transport::Channel;

pub const TIMEOUT: Duration = Duration::from_secs(60);

pub async fn wait_for_cosmos_online(contact: &Contact, timeout: Duration) {
    let start = Instant::now();
    while let Err(CosmosGrpcError::NodeNotSynced) | Err(CosmosGrpcError::ChainNotRunning) =
        contact.wait_for_next_block(timeout).await
    {
        sleep(Duration::from_secs(1)).await;
        if Instant::now() - start > timeout {
            panic!("Cosmos node has not come online during timeout!")
        }
    }
    contact.wait_for_next_block(timeout).await.unwrap();
    contact.wait_for_next_block(timeout).await.unwrap();
    contact.wait_for_next_block(timeout).await.unwrap();
}

/// gets the Cosmos last event nonce, no matter how long it takes.
pub async fn get_last_event_nonce_with_retry(
    client: &mut GravityQueryClient<Channel>,
    our_cosmos_address: CosmosAddress,
    prefix: String,
) -> u64 {
    let mut res =
        get_last_event_nonce_for_validator(client, our_cosmos_address, prefix.clone()).await;
    while res.is_err() {
        error!(
            "Failed to get last event nonce, is the Cosmos GRPC working? {:?}",
            res
        );
        sleep(RETRY_TIME).await;
        res = get_last_event_nonce_for_validator(client, our_cosmos_address, prefix.clone()).await;
    }
    res.unwrap()
}

pub async fn wait_for_tx_with_retry(contact: &Contact, response: &TxResponse) -> Result<TxResponse, CosmosGrpcError> {
    let mut res = contact.wait_for_tx(response.clone(), TIMEOUT).await;

    let mut counter: i32 = 0;
    while res.is_err() {
        info!("Wait for tx at iteration {} of 12", counter);
        sleep(RETRY_TIME).await;
        res = contact.wait_for_tx(response.clone(), TIMEOUT).await;
        counter += 1;

        if counter == 12 { // wait for 1 minute (12 * 5 = 60 seconds)
            break;
        }
    }
    
    return res;
}