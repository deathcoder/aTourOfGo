# A Tour of Go
https://tour.golang.org/

#### well this is taking longer than i anticipated, so will add things here when i encounter use cases

## Progress

### - Basics [9/17]
### - Flow control [0/14]
### - More types [0/27]
### - Methods [0/26]
### - Concurrency [0/11]

## Run all exercises

```bash
./run all
```

## Run all exercises and display documentation
```bash
./run all-doc
```

## Run exercise
```bash
./run <exercise-name>
```

### Setup new exercise
The following script setups a project under `/cmd` folder, 
```bash
./new <exercise-name>
```

### project setup inspiration
https://www.wolfe.id.au/2020/03/10/starting-a-go-project/

https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/

but with a few tweaks:

## /cmd

This folder contains the main entry point files for the projects

## /lib
This package holds the private library code used in your services, it is specific to the function of the service and not shared with other services. 
One thing to note is this privacy is enforced by the compiler itself, see the Go 1.4 release notes for more details.

## /interface
This folder contains code which is OK for other services to consume, this may include API clients, or utility functions which may be handy for other projects but don’t justify their own project.
Personally I prefer to use this over internal, mainly as I like to keep things open for reuse in most of projects.

