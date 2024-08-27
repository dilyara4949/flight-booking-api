# Flight Booking API - Simple Event-Driven Architecture (Sequence Diagram)

```mermaid
sequenceDiagram
    User->>UserHandler: Sign Up
    alt Failed
        UserHandler--xUser: Sign Up failed
    else OK
        UserHandler->>Kafka: User signed up
        Kafka->>NotifyService: Send Welcome Email
        NotifyService-->>User: Welcome Email Sent
    end
```