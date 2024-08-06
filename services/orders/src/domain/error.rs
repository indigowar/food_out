use core::fmt;

#[derive(Debug, Clone)]
pub enum ValidationError {
    Single { field: String, message: String },
    Multiple(Vec<Self>),
}

impl fmt::Display for ValidationError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::Single { field, message } => {
                write!(f, "Invalid field {}: {}", field, message)
            }
            Self::Multiple(errors) => {
                write!(f, "Multiple errors occurred:")?;
                for e in errors {
                    write!(f, "\n - {}", e)?;
                }
                Ok(())
            }
        }
    }
}

impl ValidationError {
    pub fn single(field: String, message: String) -> ValidationError {
        ValidationError::Single { field, message }
    }

    pub fn multiple(errors: Vec<ValidationError>) -> ValidationError {
        ValidationError::Multiple(errors)
    }

    pub fn unite(left: Option<Self>, right: Option<Self>) -> Option<Self> {
        match (left, right) {
            (Some(l), Some(r)) => Some(Self::flatten(l, r)),
            (Some(l), None) => Some(l),
            (None, Some(r)) => Some(r),
            (None, None) => None,
        }
    }

    pub fn flatten(left: Self, right: Self) -> Self {
        match (left.clone(), right.clone()) {
            (Self::Single { .. }, Self::Single { .. }) => Self::multiple(vec![left, right]),
            (Self::Single { .. }, Self::Multiple(mut e)) => {
                e.push(left);
                Self::multiple(e)
            }
            (ValidationError::Multiple(mut e), ValidationError::Single { .. }) => {
                e.push(right);
                Self::multiple(e)
            }
            (ValidationError::Multiple(mut l), ValidationError::Multiple(r)) => {
                l.extend(r.into_iter());
                Self::multiple(l)
            }
        }
    }
}
