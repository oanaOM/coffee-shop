# GopherCon UK -2023

## Scaling Coffee with Goroutines

A RESTful API to serves coffees, powered by goroutines.

This project was done during GopherCON UK 2023, during Sadie's Freeman (@sadief) workshop.

### Requirements

- it has to be scalable and fast!

### Implementation steps:

- Server coffee without concurrency
- Server more than 1 coffee
- Add concurrency: use goroutines
- Add WaitGroup:
  Why? to make sure that all goroutines we start will also finish
- Add more goroutines: so we can run each function concurrent

### How to run it?

- in root folder run `make run` to start the docker container
- navigate to `http://localhost:8080/serve-coffee/50000` to hit the endpoint and try making MANY coffee. Checkout the logs to see the time it took to make them. In my case, 50000 coffee took 13.390567749s
- in root folder run `make down` to stop the docker container and remove image.

### What's next?

- Scale vertically by adding more pods

### Testing Results

Configuration 1

```
resources:
            limits:
              cpu: 1
              memory: 3092M
            requests:
              cpu: 1
              memory: 1024M
```

// Took 5.057038683s to serve coffee, customer no: 50000

Configuration 2

```
          resources:
            limits:
              cpu: 4
              memory: 3092M
            requests:
              cpu: 2
              memory: 1024M
```

Took 9.362377223s to serve coffee, customer no: 50000

Why? 🤔

##### Resources

https://github.com/sadief/gophercon-2023-slides-code/tree/main
