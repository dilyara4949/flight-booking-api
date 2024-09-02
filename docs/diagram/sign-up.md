# Flight Booking API - Sign Up

```mermaid
sequenceDiagram
    User->>UserHandler: Sign Up
    alt OK
        UserHandler->>Kafka: User signed up
        Kafka->>NotifyService: Send Welcome Email
        NotifyService-->>User: Welcome Email Sent
    else Failed
        UserHandler--xUser: Sign Up failed
    end
```
