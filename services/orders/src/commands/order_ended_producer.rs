use async_trait::async_trait;

use crate::domain::Order;

#[async_trait]
pub trait OrderEndedProducer {
    async fn produce(&self, order: &Order) -> Result<(), String>;
}
