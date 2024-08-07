use chrono::{DateTime, Local};
use uuid::Uuid;

use super::order_storage::{GetError, OrderStorage, UpdateError};
use crate::domain::Order;

pub struct Command {
    storage: Box<dyn OrderStorage>,
}

impl Command {
    pub fn new(storage: Box<dyn OrderStorage>) -> Self {
        Self { storage }
    }

    pub async fn take_order(
        &mut self,
        id: Uuid,
        courier: Uuid,
        taken_at: DateTime<Local>,
    ) -> Result<(), Error> {
        let mut order = self.get(id).await?;
        match order.mark_taken(courier, taken_at) {
            Err(_) => Err(Error::AlreadyTaken),
            Ok(_) => self.save(&order).await,
        }
    }

    async fn get(&self, id: Uuid) -> Result<Order, Error> {
        match self.storage.get(id).await {
            Ok(o) => Ok(o),
            Err(e) => {
                log::warn!("take_order: {}", e);
                match e {
                    GetError::NotFound(_) => Err(Error::NotFound),
                    GetError::Unexpected(_) => Err(Error::Internal),
                }
            }
        }
    }

    async fn save(&mut self, order: &Order) -> Result<(), Error> {
        match self.storage.update(&order).await {
            Ok(_) => Ok(()),
            Err(e) => {
                log::warn!("take_order: {}", e);
                match e {
                    UpdateError::NotFound(_) => Err(Error::NotFound),
                    UpdateError::AlreadyExists(_) => Err(Error::Internal),
                    UpdateError::Unexpected(_) => Err(Error::Internal),
                }
            }
        }
    }
}

#[derive(Debug)]
pub enum Error {
    NotFound,
    AlreadyTaken,
    Internal,
}
