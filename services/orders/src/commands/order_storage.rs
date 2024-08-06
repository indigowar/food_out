use core::fmt;

use async_trait::async_trait;
use uuid::Uuid;

use crate::domain::Order;

#[async_trait]
pub trait OrderStorage {
    async fn get(&self, id: Uuid) -> Result<Order, GetError>;
    async fn add(&mut self, order: &Order) -> Result<(), AddError>;
    async fn delete(&mut self, id: Uuid) -> Result<(), DeleteError>;
    async fn update(&mut self, order: &Order) -> Result<(), UpdateError>;
}

#[derive(Debug)]
pub struct Error {
    pub object: String,
    pub field: String,
    pub message: String,
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "object: {}, field: {}, message: {}",
            self.object, self.field, self.message
        )
    }
}

#[derive(Debug)]
pub enum GetError {
    NotFound(Error),
    Unexpected(String),
}

impl fmt::Display for GetError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            GetError::NotFound(e) => write!(f, "OrderStorage.get: not found: {}", e),
            GetError::Unexpected(e) => write!(f, "OrderStorage.get: unexpected: {}", e),
        }
    }
}

impl GetError {
    pub fn not_found(object: String, field: String, message: String) -> Self {
        Self::NotFound(Error {
            object,
            field,
            message,
        })
    }

    pub fn unexpected(message: String) -> Self {
        Self::Unexpected(message)
    }
}

#[derive(Debug)]
pub enum AddError {
    AlreadyExists(Error),
    Unexpected(String),
}

impl fmt::Display for AddError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            AddError::AlreadyExists(e) => write!(f, "OrderStorage.add: already exists: {}", e),
            AddError::Unexpected(e) => write!(f, "OrderStorage.add: unexpected: {}", e),
        }
    }
}

impl AddError {
    pub fn already_exists(object: String, field: String, message: String) -> Self {
        Self::AlreadyExists(Error {
            object,
            field,
            message,
        })
    }

    pub fn unexpected(message: String) -> Self {
        Self::Unexpected(message)
    }
}

#[derive(Debug)]
pub enum DeleteError {
    NotFound(Error),
    Unexpected(String),
}

impl fmt::Display for DeleteError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            DeleteError::NotFound(e) => write!(f, "OrderStorage.delete: not found: {}", e),
            DeleteError::Unexpected(e) => write!(f, "OrderStorage.delete: unexpected: {}", e),
        }
    }
}

impl DeleteError {
    pub fn not_found(object: String, field: String, message: String) -> Self {
        Self::NotFound(Error {
            object,
            field,
            message,
        })
    }

    pub fn unexpected(message: String) -> Self {
        Self::Unexpected(message)
    }
}

#[derive(Debug)]
pub enum UpdateError {
    NotFound(Error),
    AlreadyExists(Error),
    Unexpected(String),
}

impl fmt::Display for UpdateError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            UpdateError::NotFound(e) => write!(f, "OrderStorage.delete: not found: {}", e),
            UpdateError::AlreadyExists(e) => {
                write!(f, "OrderStorage.update: already exists: {}", e)
            }
            UpdateError::Unexpected(e) => write!(f, "OrderStorage.delete: unexpected: {}", e),
        }
    }
}

impl UpdateError {
    pub fn not_found(object: String, field: String, message: String) -> Self {
        Self::NotFound(Error {
            object,
            field,
            message,
        })
    }

    pub fn already_exists(object: String, field: String, message: String) -> Self {
        Self::AlreadyExists(Error {
            object,
            field,
            message,
        })
    }

    pub fn unexpected(message: String) -> Self {
        Self::Unexpected(message)
    }
}
