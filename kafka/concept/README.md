### Kafka concept

Kafka combines three key capabilities so you can implement your use cases for event streaming end-to-end with a single battle-tested solution:

1. To **publish (write)** and **subscribe to (read)** streams of events, including continuous import/export of your data from other systems.
2. To **store streams of events durably and reliably** for as long as you want.
3. To process **streams of events** as they occur or retrospectively.

Kafka Compoment:

1. **Event**: (record or message)
2. **Topic**: store event.
3. **Partition**: Topic is spread over a number of "buckets" located on different Kafka brokers.
4. **Broker**:
5. **Producer**: publish event to Kafka.
6. **Consumer**: subcribe event.