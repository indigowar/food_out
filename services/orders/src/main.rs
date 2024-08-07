use chrono::{Duration, Local};
use uuid::Uuid;

mod commands;
mod domain;
mod storage;

#[tokio::main]
async fn main() {
    let order = domain::Order::create(
        Uuid::new_v4(),
        "address".to_string(),
        Uuid::new_v4(),
        vec![],
        Local::now() + Duration::days(15),
    );

    match order {
        Err(e) => eprintln!("{}", e),
        Ok(o) => println!(
            "Order {} of {} for {} created",
            o.id, o.customer, o.restaurant
        ),
    }
}
