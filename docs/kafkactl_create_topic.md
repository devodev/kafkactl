## kafkactl create topic

Topic API

```
kafkactl create topic TOPIC_NAME [flags]
```

### Options

```
      --config stringArray       Topic config (key=value)
  -h, --help                     help for topic
      --partitions-count int     Topic partitions count
      --replication-factor int   Topic replication factor
```

### Options inherited from parent commands

```
  -f, --config-file string   Configuration file path
  -H, --header stringArray   Additional HTTP header(s)
  -v, --log-level string     Log level (debug, info, warn, error, fatal, panic) (default "warning")
  -o, --output string        Output format (table,template,json) (default "table")
```

### SEE ALSO

* [kafkactl create](kafkactl_create.md)	 - Create resources

