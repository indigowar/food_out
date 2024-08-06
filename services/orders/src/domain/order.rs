use core::fmt;

use chrono::{DateTime, Local};
use uuid::Uuid;

use super::info::*;
use super::ValidationError;

#[derive(Default)]
pub struct Order {
    pub id: Uuid,
    pub customer: Uuid,
    pub customer_address: String,
    pub restaurant: Uuid,
    pub created_at: DateTime<Local>,

    pub products: Vec<Product>,

    pub accepted: Option<AcceptanceInfo>,
    pub courier: Option<CourierInfo>,
    pub payment: Option<PaymentInfo>,
    pub cancellation: Option<CancellationInfo>,

    pub cooking_started_at: Option<DateTime<Local>>,
    pub delivery_started_at: Option<DateTime<Local>>,
    pub delivery_completed_at: Option<DateTime<Local>>,
}

impl Order {
    pub fn create(
        customer: Uuid,
        customer_address: String,
        restaurant: Uuid,
        products: Vec<Product>,
        created_at: DateTime<Local>,
    ) -> Result<Order, ValidationError> {
        let mut order = Order {
            id: Uuid::new_v4(),
            customer,
            customer_address,
            restaurant,

            ..Self::default()
        };

        let mut actions: Vec<Box<dyn FnMut(&mut Order) -> Result<(), ValidationError>>> = vec![
            Box::new(|o| o.set_products(products.clone())),
            Box::new(|o| o.set_created_at(created_at)),
        ];

        let error = actions
            .iter_mut()
            .map(|c| match c(&mut order) {
                Ok(()) => None,
                Err(e) => Some(e),
            })
            .fold(None, ValidationError::unite);

        match error {
            None => Ok(order),
            Some(e) => Err(e),
        }
    }

    pub fn mark_accepted(
        &mut self,
        manager: Uuid,
        accepted_at: DateTime<Local>,
    ) -> Result<(), String> {
        match self.accepted {
            Some(_) => Err("order already accepted".into()),
            None => {
                self.accepted = Some(AcceptanceInfo {
                    manager,
                    accepted_at,
                });

                Ok(())
            }
        }
    }

    pub fn mark_taken(&mut self, courier: Uuid, taken_at: DateTime<Local>) -> Result<(), String> {
        match self.courier {
            Some(_) => Err("order already taken".into()),
            None => {
                self.courier = Some(CourierInfo { courier, taken_at });
                Ok(())
            }
        }
    }

    pub fn mark_payed(
        &mut self,
        transaction: String,
        payed_at: DateTime<Local>,
    ) -> Result<(), String> {
        match self.payment {
            Some(_) => Err("order already has been payed for".into()),
            None => {
                self.payment = Some(PaymentInfo {
                    transaction,
                    payed_at,
                });
                Ok(())
            }
        }
    }

    pub fn mark_cooking_started(&mut self, started_at: DateTime<Local>) -> Result<(), String> {
        match self.cooking_started_at {
            Some(_) => Err("order already received the signal".into()),
            None => {
                self.cooking_started_at = Some(started_at);
                Ok(())
            }
        }
    }

    pub fn mark_delivery_started(&mut self, started_at: DateTime<Local>) -> Result<(), String> {
        match self.delivery_started_at {
            Some(_) => Err("order already received the signal".into()),
            None => {
                self.delivery_started_at = Some(started_at);
                Ok(())
            }
        }
    }

    pub fn cancel(&mut self, id: Uuid, cancelled_at: DateTime<Local>) -> Result<(), String> {
        match self.cancellation {
            Some(_) => Err("order already has been cancelled".into()),
            None => {
                self.cancellation = Some(CancellationInfo {
                    canceller: id,
                    cancelled_at,
                });
                Ok(())
            }
        }
    }

    pub fn deliver(&mut self, timestamp: DateTime<Local>) -> Result<(), String> {
        match self.cancellation {
            Some(_) => Err("order already has been delivered".into()),
            None => {
                self.delivery_completed_at = Some(timestamp);
                Ok(())
            }
        }
    }
}

impl fmt::Display for Order {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "Order {} by customer {} for restaurant {}",
            self.id, self.customer, self.restaurant
        )
    }
}

impl fmt::Debug for Order {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "[ Order: {}, Customer: {}, Restaurant: {}",
            self.id, self.customer, self.restaurant
        )?;
        write!(
            f,
            ", Created At: {}, Products: Vec({}) ]",
            self.created_at,
            self.products.len()
        )
    }
}

impl Order {
    fn set_products(&mut self, products: Vec<Product>) -> Result<(), ValidationError> {
        if products.is_empty() {
            return Err(ValidationError::single(
                "products".into(),
                "value is empty".into(),
            ));
        }
        if !products.iter().all(|p| p.restaurant == self.restaurant) {
            return Err(ValidationError::single(
                "products".into(),
                "not all products belong to the restaurant".into(),
            ));
        }
        self.products = products;
        Ok(())
    }

    fn set_created_at(&mut self, created_at: DateTime<Local>) -> Result<(), ValidationError> {
        if created_at > Local::now() {
            return Err(ValidationError::single(
                "created_at".into(),
                "this datetime has not happened yet".into(),
            ));
        }
        self.created_at = created_at;
        Ok(())
    }
}
