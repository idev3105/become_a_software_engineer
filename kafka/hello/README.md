### Getting started

Start Kafka by docker:
``` bash
$ docker compose up
```

Attach container shell:
```bash
$ docker compose exec -it kafka /bin/bash
```

Move to Kafka home:
```bash
$ cd /opt/kafka
`````

Create a topic:
```bash
$ bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092
```
```
Created topic quickstart-events.
```

Describe topic:
```bash
$ bin/kafka-topics.sh --describe --topic quickstart-events --bootstrap-server localhost:9092
```
```bash
Topic: quickstart-events        TopicId: RugbFhgcQ2mXphhXsnkEiA PartitionCount: 1       ReplicationFactor: 1      Configs: segment.bytes=1073741824
        Topic: quickstart-events        Partition: 0    Leader: 1       Replicas: 1     Isr: 1  Elr:    LastKnownElr:
```

Send event to topic by producer:
```bash
$ bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092
>This is my first event
>This is my second even
```

Receive event by consumer:
```bash
$ bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092
This is my first event
This is my second event
```