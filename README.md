# Notification Server

Simple notification server example with go.

## Run the Application

To set up the application you need to run `zookeper` and `kafka`. For `zookeper` and `kafka` you can refer to the website [Apache Quickstart](https://kafka.apache.org/quickstart). 

After you ran the kafka and zookeper you can run the server package. The following command will automatically run the server application in your `localhost:8888`.


```bash
go run server
```


Then open another terminal and run the client to test application behavior.

```bash
go run client/client.go
```

Currently messaging queue is not connected with the notification mechanism because there is no producer set up yet. But to test if it is working you can run a producer from the terminal and it will listen the sarama topic and will log it.