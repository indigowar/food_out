use chrono::{DateTime, Local};
use uuid::Uuid;

use crate::domain::Order;

use super::{
    order_ended_producer::OrderEndedProducer,
    order_storage::{DeleteError, GetError, OrderStorage},
};

pub struct Command {
    storage: Box<dyn OrderStorage>,
    order_ended: Box<dyn OrderEndedProducer>,
}

#[derive(Debug)]
pub enum Error {
    NotFound,
    Invalid(String),
    Internal,
}

impl Command {
    pub fn new(storage: Box<dyn OrderStorage>, order_ended: Box<dyn OrderEndedProducer>) -> Self {
        Self {
            storage,
            order_ended,
        }
    }

    pub async fn cancel_order(
        &mut self,
        id: Uuid,
        canceller: Uuid,
        timestamp: DateTime<Local>,
    ) -> Result<(), Error> {
        let mut order = self.get(id).await?;
        if let Err(e) = order.cancel(canceller, timestamp) {
            return Err(Error::Invalid(e));
        }
        self.produce_event(&order).await?;
        self.delete(order.id).await
    }

    async fn get(&self, id: Uuid) -> Result<Order, Error> {
        match self.storage.get(id).await {
            Ok(o) => Ok(o),
            Err(e) => {
                log::warn!("cancel_order: {}", e);
                match e {
                    GetError::NotFound(_) => Err(Error::NotFound),
                    GetError::Unexpected(_) => Err(Error::Internal),
                }
            }
        }
    }

    async fn delete(&mut self, id: Uuid) -> Result<(), Error> {
        match self.storage.delete(id).await {
            Ok(_) => Ok(()),
            Err(e) => {
                log::warn!("cancel_order: {}", e);
                match e {
                    DeleteError::NotFound(_) => Err(Error::NotFound),
                    DeleteError::Unexpected(_) => Err(Error::Internal),
                }
            }
        }
    }

    async fn produce_event(&self, order: &Order) -> Result<(), Error> {
        match self.order_ended.produce(order).await {
            Ok(_) => Ok(()),
            Err(e) => {
                log::warn!("cancel_order: failed to produce event: {}", e);
                Err(Error::Internal)
            }
        }
    }
}
