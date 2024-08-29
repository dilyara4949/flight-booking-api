# Flight Booking API - Ticket Booking

```mermaid
sequenceDiagram
    User->>TicketHandler: Book Ticket
    alt OK
        TicketHandler->>Kafka: Send Booking Confirmation
        Kafka->>NotifyService: Send Booking Confirmation Email
        NotifyService-->>User: Booking Confirmation Email Sent
    else Failed
        TicketHandler--xUser: Booking failed
    end
```
