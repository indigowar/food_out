use chrono::{DateTime, Local};
use uuid::Uuid;

use crate::domain::{Order, Product, ValidationError};

use super::order_storage::OrderStorage;

pub struct Command {
    order_storage: Box<dyn OrderStorage>,
}

#[derive(Debug)]
pub enum Error {
    InvalidData(ValidationError),

    AlreadyExists,
    Internal,
}

impl Command {
    pub fn new(order_storage: Box<dyn OrderStorage>) -> Self {
        Self { order_storage }
    }

    pub async fn create_order(
        &mut self,
        customer: Uuid,
        customer_address: String,
        restaurant: Uuid,
        products: Vec<Product>,
        timestamp: DateTime<Local>,
    ) -> Result<(), Error> {
        let order = Order::create(customer, customer_address, restaurant, products, timestamp);
        if order.is_err() {
            return Err(Error::InvalidData(order.unwrap_err()));
        }
        let order = order.unwrap();

        match self.order_storage.add(&order).await {
            Ok(()) => Ok(()),
            Err(e) => {
                log::warn!("create_order: {}", e);

                Err(match e {
                    super::order_storage::AddError::AlreadyExists(_) => Error::AlreadyExists,
                    super::order_storage::AddError::Unexpected(_) => Error::Internal,
                })
            }
        }
    }
}
