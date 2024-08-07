use uuid::Uuid;

use crate::domain;

struct Order {
    pub id: Uuid,
}

struct Product {
    pub id: Uuid,
    pub order: Uuid,
    pub restaurant: Uuid,
    pub name: String,
    pub picture: String,
    pub price: f64,
    pub categories: Vec<String>,
}

impl Into<domain::Product> for Product {
    fn into(self) -> domain::Product {
        domain::Product {
            id: self.id,
            restaurant: self.restaurant,
            name: self.name,
            picture: self.picture,
            price: self.price,
            categories: self.categories,
        }
    }
}

impl From<domain::Product> for Product {
    fn from(value: domain::Product) -> Self {
        Self {
            id: value.id,
            order: Uuid::default(),
            restaurant: value.restaurant,
            name: value.name,
            picture: value.picture,
            price: value.price,
            categories: value.categories,
        }
    }
}
