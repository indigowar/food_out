use std::sync::Arc;

use async_trait::async_trait;
use sqlx::PgPool;
use uuid::Uuid;

use crate::{
    commands::order_storage::{AddError, DeleteError, GetError, OrderStorage, UpdateError},
    domain::Order,
};

#[derive(Clone)]
pub struct PostgresOrderStorage {
    pool: Arc<PgPool>,
}

impl PostgresOrderStorage {
    pub fn new(pool: Arc<PgPool>) -> Self {
        Self { pool }
    }
}

#[async_trait]
impl OrderStorage for PostgresOrderStorage {
    async fn get(&self, id: Uuid) -> Result<Order, GetError> {
        todo!()
    }

    async fn add(&mut self, order: &Order) -> Result<(), AddError> {
        todo!()
    }

    async fn delete(&mut self, id: Uuid) -> Result<(), DeleteError> {
        todo!()
    }

    async fn update(&mut self, order: &Order) -> Result<(), UpdateError> {
        todo!()
    }
}

impl PostgresOrderStorage {}
