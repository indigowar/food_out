use chrono::{DateTime, Local};
use uuid::Uuid;

use super::ValidationError;

#[derive(Debug, Clone)]
pub struct Product {
    pub id: Uuid,
    pub restaurant: Uuid,
    pub name: String,
    pub picture: String,
    pub price: f64,
    pub categories: Vec<String>,
}

impl Product {
    pub fn new(
        id: Uuid,
        restaurant: Uuid,
        name: String,
        picture: String,
        price: f64,
        categories: Vec<String>,
    ) -> Result<Self, ValidationError> {
        if price < 0.0 {
            return Err(ValidationError::single(
                "price".into(),
                "should be above zero".into(),
            ));
        }

        Ok(Self {
            id,
            restaurant,
            name,
            picture,
            price,
            categories,
        })
    }
}

#[derive(Debug, Clone)]
pub struct AcceptanceInfo {
    pub manager: Uuid,
    pub accepted_at: DateTime<Local>,
}

#[derive(Debug, Clone)]
pub struct CourierInfo {
    pub courier: Uuid,
    pub taken_at: DateTime<Local>,
}

#[derive(Debug, Clone)]
pub struct PaymentInfo {
    pub transaction: String,
    pub payed_at: DateTime<Local>,
}

#[derive(Debug, Clone)]
pub struct CancellationInfo {
    pub canceller: Uuid,
    pub cancelled_at: DateTime<Local>,
}
