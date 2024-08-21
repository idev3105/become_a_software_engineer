# Producers

- Producers write data to topics (which are made of partitions).
- Producers know to which partition to write to (and which Kafka broker has it).
- In case of Kafka broker failures, Producers will automatically recover.

## Message keys

- Producer can choose o send a **key** with the message (string, number, binary, etc,...)
- Key is created by producer.
- If **key = null**, data is send round robin (partition 0 -> partition 1 -> partition 3 -> ...)
- If **key != null**, then all messages for that key will always go to the same partition (hashing)
- A key are typically send if you need message ordering for a specific field (ex: truck_id)
